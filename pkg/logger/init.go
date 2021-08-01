package logger

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func Init() {
	logger = NewLogger("")
	if logger == nil {
		panic("创建日志失败")
	}

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()/zap.S()调用即可
	zap.ReplaceGlobals(logger)
}
