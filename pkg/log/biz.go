package log

import (
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

func CopierErrLog(l *zap.Logger, err error) {
	BizLog(l, copierErrTitle, err)
}

func TransformErrLog(l *zap.Logger, err error) {
	BizLog(l, transformErrTitle, err)
}

func JsonMarshalErrLog(l *zap.Logger, err error) {
	BizLog(l, jsonMarshalErrTitle, err)
}

func CacheErrLog(l *zap.Logger, err error) {
	BizLog(l, cacheErrTitle, err)
}

func MysqlErrLog(l *zap.Logger, err error) {
	BizLog(l, mysqlErrTitle, err)
}

func BizLog(l *zap.Logger, title string, err error) {
	l.Sugar().Infof("%s:%s", title, err)
}
