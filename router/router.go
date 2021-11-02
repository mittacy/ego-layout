package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/pkg/log"
	"time"
)

func InitRouter(r *gin.Engine) {
	// 记录请求日志
	requestLog := log.New("request")
	r.Use(ginzap.Ginzap(requestLog.GetZap(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(requestLog.GetZap(), true))

	globalPath := "/api"

	g := r.Group(globalPath)
	{
		g.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"data": "", "msg": "success"})
		})
	}
}
