package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"io/ioutil"
	"net/http"
)

func main() {
	// tracer配置信息
	cfg := jaegercfg.Configuration{
		ServiceName: "helloJaeger",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: "http://192.168.163.131:14268/api/traces",
		},
	}

	// 创建tracer对象
	jLogger := jaegerlog.StdLogger
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jLogger))
	defer closer.Close()
	if err != nil {
		return
	}

	// 创建第一个span main
	parentSpan := tracer.StartSpan("main")
	defer parentSpan.Finish()

	Hello(tracer, parentSpan)
}

func Hello(tracer opentracing.Tracer, parentSpan opentracing.Span) {
	childSpan := tracer.StartSpan("hello", opentracing.ChildOf(parentSpan.Context()))
	defer childSpan.Finish()

	url := "http://127.0.0.1:8888/hello?username=lili"
	req, _ := http.NewRequest("GET", url, nil)

	// 对Span打标签
	ext.SpanKindRPCClient.Set(childSpan)
	ext.HTTPUrl.Set(childSpan, url)
	ext.HTTPMethod.Set(childSpan, "GET")

	// 打包当前进程上下文到header中
	tracer.Inject(childSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	resp, err := http.DefaultClient.Do(req)
	if resp == nil {
		return
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
