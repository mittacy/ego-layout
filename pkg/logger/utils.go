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
)

// CopierErrLog 结构体转化错误
// @param err
func CopierErrLog(err error) {
	LogWithStack(copierErrTitle, err)
}

// TransformErrLog 响应包装错误
// @param err
func TransformErrLog(err error) {
	LogWithStack(transformErrTitle, err)
}

// JsonMarshalErrLog json序列化与反序列化日志错误
// @param err
func JsonMarshalErrLog(err error) {
	LogWithStack(jsonMarshalErrTitle, err)
}

// CacheErrLog 缓存错误日志
// @param err
func CacheErrLog(err error) {
	LogWithStack(cacheErrTitle, err)
}

// LogWithStack 写带有调用栈的日志
// @param title 标题
// @param err 错误信息
func LogWithStack(title string, err error) {
	zap.L().Error(fmt.Sprintf("%s:%s", title, err),
		zap.String("trace", fmt.Sprintf("%+v", err)))
}
