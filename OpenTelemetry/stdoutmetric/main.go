package main

import (
	"context"
	"encoding/json"
	"go.opentelemetry.io/otel/metric/instrument"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric/unit"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

var (
	// Sat Jan 01 2000 00:00:00 GMT+0000.
	now = time.Date(2000, time.January, 01, 0, 0, 0, 0, time.FixedZone("GMT", 0))

	res = resource.NewSchemaless(
		semconv.ServiceNameKey.String("stdoutmetric-example"),
	)

	mockData = metricdata.ResourceMetrics{
		Resource: res,
		ScopeMetrics: []metricdata.ScopeMetrics{
			{
				Scope: instrumentation.Scope{Name: "example", Version: "v0.0.1"},
				Metrics: []metricdata.Metrics{
					{
						Name:        "requests",
						Description: "Number of requests received",
						Unit:        unit.Dimensionless,
						Data: metricdata.Sum[int64]{
							IsMonotonic: true,
							Temporality: metricdata.DeltaTemporality,
							DataPoints: []metricdata.DataPoint[int64]{
								{
									Attributes: attribute.NewSet(attribute.String("server", "central")),
									StartTime:  now,
									Time:       now.Add(1 * time.Second),
									Value:      5,
								},
							},
						},
					},
					{
						Name:        "latency",
						Description: "Time spend processing received requests",
						Unit:        unit.Milliseconds,
						Data: metricdata.Histogram{
							Temporality: metricdata.DeltaTemporality,
							DataPoints: []metricdata.HistogramDataPoint{
								{
									Attributes:   attribute.NewSet(attribute.String("server", "central")),
									StartTime:    now,
									Time:         now.Add(1 * time.Second),
									Count:        10,
									Bounds:       []float64{1, 5, 10},
									BucketCounts: []uint64{1, 3, 6, 0},
									Sum:          57,
								},
							},
						},
					},
					{
						Name:        "temperature",
						Description: "CPU global temperature",
						Unit:        unit.Unit("cel(1 K)"),
						Data: metricdata.Gauge[float64]{
							DataPoints: []metricdata.DataPoint[float64]{
								{
									Attributes: attribute.NewSet(attribute.String("server", "central")),
									Time:       now.Add(1 * time.Second),
									Value:      32.4,
								},
							},
						},
					},
				},
			},
		},
	}
)

func main() {
	// Print with a JSON encoder that indents with two spaces.
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	exp, err := stdoutmetric.New(stdoutmetric.WithEncoder(enc))
	if err != nil {
		panic(err)
	}

	// Register the exporter with an SDK via a periodic reader.
	sdk := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(exp)),
	)

	ctx := context.Background()
	// This is where the sdk would be used to create a Meter and from that
	// instruments that would make measurments of your code. To simulate that
	// behavior, call export directly with mocked data.
	meter := sdk.Meter("test")
	counter, err := meter.SyncFloat64().Counter("foo", instrument.WithDescription("a simple counter"))
	if err != nil {
		log.Fatal(err)
	}
	counter.Add(ctx, 1)

	_ = exp.Export(ctx, mockData)

	// Ensure the periodic reader is cleaned up by shutting down the sdk.
	_ = sdk.Shutdown(ctx)
}
