package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"net/http"
	"time"
)

var tracer = otel.Tracer("otelhttp")

func sleepy(ctx context.Context) {
	_, span := tracer.Start(ctx, "sleepy")
	defer span.End()

	sleepTime := 1 * time.Second
	time.Sleep(sleepTime)

	span.SetAttributes(attribute.Int("sleep.duration", int(sleepTime)))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! I am instrumented automatically!")
	ctx := r.Context()
	sleepy(ctx)
}

func main() {
	handler := http.HandlerFunc(httpHandler)
	wrappedHandler := otelhttp.NewHandler(handler, "hello")
	// 定义路由
	http.Handle("/hello", wrappedHandler)

	// 开启监听
	http.ListenAndServe(":9030", nil)
}
