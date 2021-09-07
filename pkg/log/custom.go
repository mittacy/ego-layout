package log

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

// 业务常用的快捷记录日志方法

func (l *Logger) CopierErrLog(err error) {
	l.LogWithStack(copierErrTitle, err)
}

func (l *Logger) TransformErrLog(err error) {
	l.LogWithStack(transformErrTitle, err)
}

func (l *Logger) JsonMarshalErrLog(err error) {
	l.LogWithStack(jsonMarshalErrTitle, err)
}

func (l *Logger) CacheErrLog(err error) {
	l.LogWithStack(cacheErrTitle, err)
}

func (l *Logger) MysqlErrLog(err error) {
	l.LogWithStack(mysqlErrTitle, err)
}

func (l *Logger) LogWithStack(title string, err error) {
	l.Error(fmt.Sprintf("%s:%s", title, err),
		zap.String("trace", fmt.Sprintf("%+v", err)))
}
