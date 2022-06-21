package main

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
)

func HelloFunc(w http.ResponseWriter, r *http.Request) {
	// 断路器配置
	hystrix.ConfigureCommand("test", hystrix.CommandConfig{
		Timeout:                1000,  // 超时时间
		MaxConcurrentRequests:  10,    // 最大并发数
		RequestVolumeThreshold: 10,    // 统计指定秒数内的请求数
		SleepWindow:            10000, // 开启断路器后指定时间内尝试关闭
		ErrorPercentThreshold:  30,    // 请求失败错误比例
	})
	// Do为同步方法，包含三个参数，第一个参数为断路器配置；第二个参数为正常执行方法；第三个参数为发生故障时执行的方法；
	// Go为异步方法，参数和Do方法一样。
	hystrix.Do("test", func() error { // run 正常逻辑
		resp, err := http.Get("https://www.baidu.com")
		if err != nil {
			return err
		}
		if resp.Body == nil {
			return errors.New("resp.body is nil")
		}
		fmt.Println("success")
		w.Write([]byte("success"))
		return nil
	}, func(err error) error { // fallback 降级处理
		w.Write([]byte(err.Error()))
		return nil
	})
}

func main() {
	http.HandleFunc("/hello", HelloFunc)
	http.ListenAndServe(":8081", nil)
}
