package main

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log"
	"net/http"
)

var tracer trace.Tracer

// 定义一个exporter导出器
func newExporter() (*jaeger.Exporter, error) {
	url := "http://10.0.98.16:14268/api/traces"
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}

// 定义一个资源
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("serviceDemo"),
			semconv.ServiceVersionKey.String("v1.0.0"),
			attribute.String("RUNMODE", "dev"),
		),
	)
	return r
}

func initTracer() (*sdktrace.TracerProvider, error) {
	exporter, err := newExporter()
	if err != nil {
		return nil, err
	}

	// 创建追踪提供对象
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(newResource()),
	)

	// 声明为全局对象
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	// 设置一个名为otelhttp2的追踪者
	tracer = tp.Tracer("otelhttp2_server")
	return tp, err
}

func test(ctx context.Context) string {
	_, span := tracer.Start(ctx, "test")
	defer span.End()
	result := "hello world"
	span.SetAttributes(attribute.String("name", result))
	return result
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := trace.SpanFromContext(ctx)

	// 从context中获取baggage
	bag := baggage.FromContext(ctx)
	span.AddEvent("handling this...", trace.WithAttributes(attribute.Key("username").String(bag.Member("username").Value())))
	str := test(ctx)
	io.WriteString(w, str)
}

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	handler := http.HandlerFunc(helloHandler)
	wrapperHandler := otelhttp.NewHandler(handler, "Hello")
	http.Handle("/hello", wrapperHandler)
	err = http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal(err)
	}
}
