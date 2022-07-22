package main

import (
	"go.uber.org/zap"
	"net/http"
)

var sugarLogger *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching url：%s ; Error = %s", url, err)
		return
	}
	sugarLogger.Infof("Success.. statusCode=%s ； url = %s", resp.Status, url)
	resp.Body.Close()
}

func main() {
	// 应用退出前刷新任何缓冲日志条目
	defer sugarLogger.Sync()

	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}
