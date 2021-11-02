package log

const (
	jsonMarshalErrTitle = "json marshal err"
	copierErrTitle      = "copier err"
	transformErrTitle   = "transform err"
	cacheErrTitle       = "cache err"
	mysqlErrTitle       = "mysql err"
)

// 业务常用的快捷记录日志方法

func CopierErrLog(l *Logger, err error) {
	BizLog(l, copierErrTitle, err)
}

func TransformErrLog(l *Logger, err error) {
	BizLog(l, transformErrTitle, err)
}

func JsonMarshalErrLog(l *Logger, err error) {
	BizLog(l, jsonMarshalErrTitle, err)
}

func CacheErrLog(l *Logger, err error) {
	BizLog(l, cacheErrTitle, err)
}

func MysqlErrLog(l *Logger, err error) {
	BizLog(l, mysqlErrTitle, err)
}

func BizLog(l *Logger, title string, err error) {
	l.Sugar().Errorf("%s: %s", title, err)
}
