package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 业务常用的快捷记录日志方法

func (l *Logger) CopierErrLogWithTrace(c *gin.Context, err error, req interface{}) {
	l.BizErrorLogWithTrace(c, copierErrTitle, err, req)
}

func (l *Logger) TransformErrLogWithTrace(c *gin.Context, err error, req interface{}) {
	l.BizErrorLogWithTrace(c, transformErrTitle, err, req)
}

func (l *Logger) JsonMarshalErrLogWithTrace(c *gin.Context, err error, req interface{}) {
	l.BizErrorLogWithTrace(c, jsonMarshalErrTitle, err, req)
}

func (l *Logger) CacheErrLogWithTrace(c *gin.Context, err error, req interface{}) {
	l.BizErrorLogWithTrace(c, cacheErrTitle, err, req)
}

func (l *Logger) MysqlErrLogWithTrace(c *gin.Context, err error, req interface{}) {
	l.BizErrorLogWithTrace(c, mysqlErrTitle, err, req)
}

func (l *Logger) UnknownErrLogWithTrace(c *gin.Context, err error, req interface{}) {
	l.BizErrorLogWithTrace(c, unknownErrTitle, err, req)
}

// BizErrorLog 业务错误日志
// @param title 信息记录
// @param err 错误，如果为github.com/pkg/errors类型包含错误堆栈，将会打印出堆栈信息
// @param req 请求数据
func (l *Logger) BizErrorLogWithTrace(c *gin.Context, title string, err error, req interface{}) {
	l.ErrorwWithTrace(c, title, "req", req, "err", err, "callstack", fmt.Sprintf("%+v", err))
}
