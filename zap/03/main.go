package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

var logger *zap.Logger

func init() {
	writeSyncer := getWriteSyncer()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
}

func getWriteSyncer() zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 5,
		Compress:   false,
	}
	return zapcore.AddSync(lumberjackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
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
