package logger

import (
	"fmt"
	"go.uber.org/zap"
)

const (
	jsonMarshalErrTitle = "json marshal err"
	copierErrTitle      = "copier err"
	transformErrTitle   = "transform err"
	cacheErrTitle       = "cache err"
	mysqlErrTitle       = "mysql err"
)

type CustomLogger struct {
	*zap.Logger
}

func NewCustomLogger(name string) *CustomLogger {
	return &CustomLogger{
		NewLogger(name),
	}
}

// CopierErrLog 结构体转化错误
// @param err
func (logger *CustomLogger) CopierErrLog(err error) {
	logger.LogWithStack(copierErrTitle, err)
}

// TransformErrLog 响应包装错误
// @param err
func (logger *CustomLogger) TransformErrLog(err error) {
	logger.LogWithStack(transformErrTitle, err)
}

// JsonMarshalErrLog json序列化与反序列化日志错误
// @param err
func (logger *CustomLogger) JsonMarshalErrLog(err error) {
	logger.LogWithStack(jsonMarshalErrTitle, err)
}

// CacheErrLog 缓存错误日志
// @param err
func (logger *CustomLogger) CacheErrLog(err error) {
	logger.LogWithStack(cacheErrTitle, err)
}

func (logger *CustomLogger) MysqlErrLog(err error) {
	logger.LogWithStack(mysqlErrTitle, err)
}

// LogWithStack 写带有调用栈的日志
// @param title 标题
// @param err 错误信息
func (logger *CustomLogger) LogWithStack(title string, err error) {
	logger.Error(fmt.Sprintf("%s:%s", title, err),
		zap.String("trace", fmt.Sprintf("%+v", err)))
}
