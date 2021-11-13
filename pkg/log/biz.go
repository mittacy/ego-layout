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
	unknownErrTitle		= "unknown err"
)

// 业务常用的快捷记录日志方法

func (l *Logger) CopierErrLog(err error) {
	l.BizErrorLog(copierErrTitle, err)
}

func (l *Logger) TransformErrLog(err error) {
	l.BizErrorLog(transformErrTitle, err)
}

func (l *Logger) JsonMarshalErrLog(err error) {
	l.BizErrorLog(jsonMarshalErrTitle, err)
}

func (l *Logger) CacheErrLog(err error) {
	l.BizErrorLog(cacheErrTitle, err)
}

func (l *Logger) MysqlErrLog(err error) {
	l.BizErrorLog(mysqlErrTitle, err)
}

func (l *Logger) UnknownErrLog(err error) {
	l.BizErrorLog(unknownErrTitle, err)
}

// BizErrorLog 业务错误日志
// @param title 信息记录
// @param err 错误，如果为github.com/pkg/errors类型包含错误堆栈，将会打印出堆栈信息
func (l *Logger) BizErrorLog(title string, err error) {
	msg := fmt.Sprintf("%s: %s", title, err)
	l.Errorw(msg, "callstack", fmt.Sprintf("%+v", err))
}
