package log

const (
	jsonMarshalErrTitle = "json marshal err"
	copierErrTitle      = "copier err"
	transformErrTitle   = "transform err"
	cacheErrTitle       = "cache err"
	mysqlErrTitle       = "mysql err"
)

// 业务常用的快捷记录日志方法

func (l *Logger) CopierErrLog(err error) {
	l.BizLog(copierErrTitle, err)
}

func (l *Logger) TransformErrLog(err error) {
	l.BizLog(transformErrTitle, err)
}

func (l *Logger) JsonMarshalErrLog(err error) {
	l.BizLog(jsonMarshalErrTitle, err)
}

func (l *Logger) CacheErrLog(err error) {
	l.BizLog(cacheErrTitle, err)
}

func (l *Logger) MysqlErrLog(err error) {
	l.BizLog(mysqlErrTitle, err)
}

func (l *Logger) BizLog(title string, err error) {
	l.Sugar().Errorf("%s: %s", title, err)
}
