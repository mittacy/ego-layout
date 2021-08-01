package response

import (
	"errors"
	"github.com/mittacy/ego-layout/pkg/checker"
	"github.com/mittacy/ego-layout/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

// TransformErrAndLog 响应包装错误响应，同时会自动记录日志
// @param c
// @param err
func TransformErrAndLog(c *gin.Context, err error) {
	logger.TransformErrLog(err)
	Unknown(c)
}

// CopierErrAndLog copier结构体转化错误响应，同时会自动记录日志
// @param c
// @param err
func CopierErrAndLog(c *gin.Context, err error) {
	logger.CopierErrLog(err)
	Unknown(c)
}

// JsonMarshalErrAndLog json序列化错误响应，同时会自动记录错误日志
// @param c
// @param err
func JsonMarshalErrAndLog(c *gin.Context, err error) {
	logger.JsonMarshalErrLog(err)
	Unknown(c)
}

// CheckErrAndLog 检查是否为指定的业务错误，记录日志并响应
// title 记录日志的标题
// sourceErr 产生的错误
// targetErr 可能的业务错误
// - 如果是这些错误，将响应错误的提示信息
// - 如果不是这些错误，将记录日志并响应未知错误
func CheckErrAndLog(c *gin.Context, title string, sourceErr error, targetErr ...error) {
	if isErr(sourceErr, targetErr...) {
		FailErr(c, sourceErr)
		return
	}

	logger.LogWithStack(title, sourceErr)
	Unknown(c)
	return
}

// ValidateErrAndLog 表单解析错误响应，记录日志并响应
// err 错误
// title 日志标记信息
func ValidateErrAndLog(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator错误
		FailMsg(c, "json错误")
		return
	}
	// validator错误进行翻译
	details := removeTopStruct(errs.Translate(checker.Trans))

	// 随机返回校验错误中的一条到 msg 字符串
	msg := ""
	for _, v := range details {
		msg = v
		break
	}

	Custom(c, http.StatusOK, 1, msg, details)
	return
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func isErr(source error, target ...error) bool {
	for _, v := range target {
		if errors.Is(source, v) {
			return true
		}
	}
	return false
}
