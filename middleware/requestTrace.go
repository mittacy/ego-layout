package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

const (
	RequestIdKey = "requestId"
)

// RequestTrace 请求追踪，添加请求id……
// @param traceIdKey 如果追踪id key能获取到追踪id，则使用该id为追踪id，否则生成新的请求id
// @return gin.HandlerFunc
func RequestTrace() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := NewRequestId()
		c.Set(RequestIdKey, requestId)
		c.Next()
	}
}

// NewRequestId 生成新的请求id
// @return string
func NewRequestId() string {
	now := time.Now()
	s := fmt.Sprintf("%s%08x%05x", "r", now.Unix(), now.UnixNano()%0x100000)

	return s + "_" + randomString(18)
}

var (
	defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

// randomString 生成随机字符串
// @param n 生成的字符串长度
// @return string 返回生成的随机字符串
func randomString(n int, randChars ...[]rune) string {
	if n <= 0 {
		return ""
	}

	var letters []rune

	if len(randChars) == 0 {
		letters = defaultLetters
	} else {
		letters = randChars[0]
	}

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}