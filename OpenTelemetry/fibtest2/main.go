package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
)

var tracer trace.Tracer

// newExporter 定义一个使用stdouttrace创建Exporter
func newExporter(w io.Writer) (sdktrace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

// newResource 定义一个分布式追踪应用的相关资源描述信息
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("fib"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}

// 定义一个斐波那契方法
func Fibonacci(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}
	if n > 93 {
		return 0, fmt.Errorf("unsupported fibonacci number %d: too large", n)
	}
	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}
	return n2 + n1, nil
}

// 定义一个应用App，该应用结构体中包含一个接收外部变量和打印日志成员变量
type App struct {
	r io.Reader
	l *log.Logger
}

// 工厂模式构造一个应用
func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{r: r, l: l}
}

// 应用执行方法，该方法调用了Poll()方法来接收外部输入的值，调用Write()方法将计算结果进行打印，内部调用Fibonacci方法
func (a *App) Run(ctx context.Context) error {
	for {
		newCtx, span := tracer.Start(ctx, "Run")

		n, err := a.Poll(newCtx)
		if err != nil {
			return err
		}
		a.Write(newCtx, n)

		span.End()
	}
}

// Poll用于接收用户输入
func (a *App) Poll(ctx context.Context) (uint, error) {
	_, span := tracer.Start(ctx, "Poll")
	defer span.End()
	a.l.Print("What Fibonacci number would you like to know: ")
	var n uint
	_, err := fmt.Fscanf(a.r, "%d\n", &n)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return 0, err
	}

	// 将输入的数字作为Span的属性
	nStr := strconv.FormatUint(uint64(n), 10)
	span.SetAttributes(attribute.String("request.n", nStr))

	return n, nil
}

// Write用于调用Fibonacci方法计算数字对应的Fibonacci数
func (a *App) Write(ctx context.Context, n uint) {
	var span trace.Span
	ctx, span = tracer.Start(ctx, "Write")
	defer span.End()
	f, err := func(ctx context.Context) (uint64, error) {
		_, span = tracer.Start(ctx, "Fibonacci")
		defer span.End()
		f, err := Fibonacci(n)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		return f, err
	}(ctx)
	if err != nil {
		a.l.Printf("Fibonacci(%d): %v\n", n, err)
	} else {
		a.l.Printf("Fibonacci(%d) = %d\n", n, f)
	}
}

func main() {
	l := log.New(os.Stdout, "", 0)

	// 创建一个文件用于存储追踪信息
	f, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	// ----------创建一个Tracer Provider------------------
	exp, err := newExporter(f)
	if err != nil {
		l.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(newResource()),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	// 注册tp为全局变量
	otel.SetTracerProvider(tp)
	tracer = tp.Tracer("ExampleService")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	app := NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}
