package log

import (
	"fmt"
)

const (
	jsonMarshalErrTitle = "json marshal err"
	copierErrTitle      = "copier err"
	transformErrTitle   = "transform err"
	cacheErrTitle       = "cache err"
	mysqlErrTitle       = "mysql err"
	unknownErrTitle     = "unknown err"
)

// 业务常用的快捷记录日志方法

func (l *Logger) CopierErrLog(err error, req interface{}) {
	l.BizErrorLog(copierErrTitle, err, req)
}

func (l *Logger) TransformErrLog(err error, req interface{}) {
	l.BizErrorLog(transformErrTitle, err, req)
}

func (l *Logger) JsonMarshalErrLog(err error, req interface{}) {
	l.BizErrorLog(jsonMarshalErrTitle, err, req)
}

func (l *Logger) CacheErrLog(err error, req interface{}) {
	l.BizErrorLog(cacheErrTitle, err, req)
}

func (l *Logger) MysqlErrLog(err error, req interface{}) {
	l.BizErrorLog(mysqlErrTitle, err, req)
}

func (l *Logger) UnknownErrLog(err error, req interface{}) {
	l.BizErrorLog(unknownErrTitle, err, req)
}

// BizErrorLog 业务错误日志
// @param title 信息记录
// @param err 错误，如果为github.com/pkg/errors类型包含错误堆栈，将会打印出堆栈信息
// @param req 请求数据
func (l *Logger) BizErrorLog(title string, err error, req interface{}) {
	l.Errorw(title, "req", req, "err", err, "callstack", fmt.Sprintf("%+v", err))
}
