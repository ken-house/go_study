package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"io"
	"net/http"
)

// CreateTracer 创建一个tracer，并继承其他进程传递过来的上下文信息
func CreateTracer(serviceName string, header http.Header) (opentracing.Tracer, opentracing.SpanContext, io.Closer, error) {
	var cfg = jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: "http://192.168.163.131:14268/api/traces",
		},
	}
	jLogger := jaegerlog.StdLogger

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jLogger))
	if err != nil {
		return nil, nil, nil, err
	}
	// 继承别的进程传递过来的上下文
	spanContext, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	return tracer, spanContext, closer, nil
}

func UseOpentracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer, spanContext, closer, _ := CreateTracer("helloWebService", c.Request.Header)
		defer closer.Close()

		// 新建一个span，依赖其他进程传递过来的span上下文
		startSpan := tracer.StartSpan(c.Request.URL.Path, ext.RPCServerOption(spanContext))
		defer startSpan.Finish()

		// 记录请求 Url
		ext.HTTPUrl.Set(startSpan, c.Request.URL.Path)
		// 记录请求方法
		ext.HTTPMethod.Set(startSpan, c.Request.Method)
		// 记录组件名称
		ext.Component.Set(startSpan, "Gin-Http")

		// 在header中加上当前进程的上下文
		c.Request = c.Request.WithContext(opentracing.ContextWithSpan(c.Request.Context(), startSpan))
		c.Next()

		// 记录当前span的状态
		ext.HTTPStatusCode.Set(startSpan, uint16(c.Writer.Status()))
	}
}

func main() {
	r := gin.Default()
	r.Use(UseOpentracing())
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	r.Run(":8888")
}
