package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/unit"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
)

// 定义一个exporter导出器
func newExporter() (*prometheus.Exporter, error) {
	return prometheus.New()
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

func newMeterProvider() *metric.MeterProvider {
	exporter, err := newExporter()
	if err != nil {
		log.Fatal(err)
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(exporter),
		metric.WithResource(newResource()),
	)
	return provider
}

func main() {
	ctx := context.Background()
	provider := newMeterProvider()
	meter := provider.Meter("prometheus_example")

	// 设置标签
	attrs := []attribute.KeyValue{
		attribute.Key("http").String("request"),
		attribute.Key("grpc").String("response"),
	}

	// This is the equivalent of prometheus.NewCounterVec
	counter, err := meter.SyncFloat64().Counter("foo", instrument.WithDescription("a simple counter"))
	if err != nil {
		log.Fatal(err)
	}
	counter.Add(ctx, 10, attribute.String("test", "req"))

	gcCount, _ := meter.AsyncInt64().Counter("gcCount")
	err = meter.RegisterCallback([]instrument.Asynchronous{gcCount}, func(ctx context.Context) {
		memStats := &runtime.MemStats{}
		// This call does work
		runtime.ReadMemStats(memStats)
		gcCount.Observe(ctx, int64(memStats.NumGC))
	})
	if err != nil {
		log.Fatal(err)
	}

	gauge, err := meter.SyncFloat64().UpDownCounter("bar", instrument.WithDescription("a fun little gauge"))
	if err != nil {
		log.Fatal(err)
	}
	gauge.Add(ctx, 100, attrs...)
	gauge.Add(ctx, -25, attrs...)
	gauge.Add(ctx, 100, attrs...)

	// 异步
	memoryUsage, err := meter.AsyncInt64().Gauge(
		"MemoryUsage",
		instrument.WithUnit(unit.Bytes),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = meter.RegisterCallback([]instrument.Asynchronous{memoryUsage}, func(ctx context.Context) {
		// 计算得到的值
		mem := 75000
		memoryUsage.Observe(ctx, int64(mem))
	})
	if err != nil {
		log.Fatal(err)
	}

	// This is the equivalent of prometheus.NewHistogramVec
	histogram, err := meter.SyncFloat64().Histogram("baz", instrument.WithDescription("a very nice histogram"))
	if err != nil {
		log.Fatal(err)
	}

	histogram.Record(ctx, 23, attrs...)
	histogram.Record(ctx, 7, attrs...)
	histogram.Record(ctx, 101, attrs...)
	histogram.Record(ctx, 105, attrs...)

	// 启动服务
	go func() {
		log.Printf("serving metrics at localhost:2223/metrics")
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":2223", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	<-ctx.Done()
}
