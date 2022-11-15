package main

import (
	"context"
	"fmt"
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
	"io/ioutil"
	"log"
	"net/http"
)

var tracer trace.Tracer

// 定义一个exporter导出器
func newExporter() (*jaeger.Exporter, error) {
	url := "http://10.0.98.16:14268/api/traces"
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}

func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("clientDemo"),
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

		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(newResource()),
	)

	// 声明为全局对象
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	// 设置一个名为otelhttp2的追踪者
	tracer = tp.Tracer("otelhttp2_client")
	return tp, err
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

	bag, _ := baggage.Parse("username=donuts")
	ctx := baggage.ContextWithBaggage(context.Background(), bag)

	ctx, span := tracer.Start(ctx, "client request", trace.WithAttributes(semconv.PeerServiceKey.String("ginServerDemo")))
	defer span.End()

	url := "http://127.0.0.1:8080/users/123"
	client := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response Received: %s\n\n\n", string(body))
}
