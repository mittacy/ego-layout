package logger

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	bizLogger *zap.Logger
	requestLogger *zap.Logger
	conf LogConf
)

func Init() {
	if err := viper.UnmarshalKey("log", &conf); err != nil {
		panic(fmt.Sprintf("获取日志配置错误: %v", err))
	}

	requestLogger = NewLogger("request")
	if requestLogger == nil {
		panic("创建请求日志失败")
	}

	initBizLogger()
}

func initBizLogger() {
	bizLogger = NewLogger("")
	if bizLogger == nil {
		panic("创建日志失败")
	}

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()/zap.S()调用即可
	zap.ReplaceGlobals(bizLogger)
}

// GetRequestLogger 获取请求日志句柄
// @return *zap.Logger
func GetRequestLogger() *zap.Logger {
	return requestLogger
}
