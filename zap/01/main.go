package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Error fetching url：", zap.String("url", url), zap.Error(err))
		return
	}
	logger.Info("Success..", zap.String("statusCode", resp.Status), zap.String("url", url))
	resp.Body.Close()
}
func main() {
	// 应用退出前刷新任何缓冲日志条目
	defer logger.Sync()

	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}
