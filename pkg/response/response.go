package response

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/apierr"
	"net/http"
)

// Custom 自定义响应
// httpCode http响应码
// apiCode 业务响应码
// msg 提示信息
// data 返回数据
func Custom(c *gin.Context, httpCode, apiCode int, msg string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": apiCode,
		"msg":  msg,
		"data": data,
	})
}

// Success 成功响应
// data 返回数据
func Success(c *gin.Context, data interface{}) {
	Custom(c, http.StatusOK, 0, "success", data)
}

// SuccessMsg 成功响应带消息提示
// data 返回数据
// msg 提示信息
func SuccessMsg(c *gin.Context, data interface{}, msg string) {
	Custom(c, http.StatusOK, 0, msg, data)
}

// Fail 失败响应
// msg 提示信息
func Fail(c *gin.Context) {
	Custom(c, http.StatusOK, 1, "fail", nil)
}

// FailMsg 带自定义信息的失败响应
// msg 自定义提示信息
func FailMsg(c *gin.Context, msg string) {
	Custom(c, http.StatusOK, 1, msg, nil)
}

// FailErr 带有错误的失败响应
// err 错误
func FailErr(c *gin.Context, err error) {
	Custom(c, http.StatusOK, apierr.ErrCode(err), err.Error(), nil)
}

// Unknown 未知错误响应
func Unknown(c *gin.Context) {
	Custom(c, http.StatusInternalServerError, 500, "unknown Error", nil)
}

// Unauthorized 未认证响应
func Unauthorized(c *gin.Context) {
	Custom(c, http.StatusUnauthorized, 401, "未认证", nil)
}

// Forbidden 权限不足响应
func Forbidden(c *gin.Context) {
	Custom(c, http.StatusForbidden, 401, "权限不足", nil)
}
