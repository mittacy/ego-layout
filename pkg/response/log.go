package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/pkg/log"
)

// CheckErrAndLog 检查是否为指定的业务错误，记录日志并响应
// title 记录日志的标题
// sourceErr 产生的错误
// targetErr 可能的业务错误
// - 如果是这些错误，将响应错误的提示信息
// - 如果不是这些错误，将记录日志并响应未知错误
func CheckErrAndLog(c *gin.Context, logger *log.Logger, req interface{}, title string, sourceErr error, targetErr ...error) {
	if isErr(sourceErr, targetErr...) {
		FailErr(c, sourceErr)
		return
	}

	logger.BizErrorLogWithTrace(c, title, sourceErr, req)
	Unknown(c)
	return
}

func isErr(source error, target ...error) bool {
	for _, v := range target {
		if errors.Is(source, v) {
			return true
		}
	}
	return false
}
