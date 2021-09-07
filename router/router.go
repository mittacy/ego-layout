package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/pkg/config"
	"github.com/mittacy/ego-layout/pkg/log"
	"time"
)

func InitRouter(r *gin.Engine) {
	// 1. 初始化控制器
	InitApi()

	// 2. 全局中间件
	requestLog := log.New("request")
	r.Use(ginzap.Ginzap(requestLog.GetZap(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(requestLog.GetZap(), true))
	//r.Use(middleware.CorsMiddleware())

	// 3. 初始化路由
	relativePath := "/api/" + config.ServerConfig.Version
	g := r.Group(relativePath) // 统一前缀
	{
		g.POST("/user", userApi.Create)
		g.DELETE("/user", userApi.Delete)
		g.PUT("/user", userApi.Update)
		g.GET("/user", userApi.Get)
		g.GET("/users", userApi.List)
	}
}
