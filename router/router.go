package router

import (
	"github.com/mittacy/ego-layout/pkg/config"
	"github.com/mittacy/ego-layout/pkg/store/cache"
	"github.com/mittacy/ego-layout/pkg/store/db"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// 1. 初始化控制器
	userApi := InitUserApi(db.GetGormDB("MYSQLKEY"), cache.GetRedisPool("REDISKEY"))

	// 2. 全局中间件
	r.Use(gin.Recovery())
	//r.Use(middleware.CorsMiddleware())

	// 3. 初始化路由
	relativePath := "/api/" + config.Global.Server.Version
	g := r.Group(relativePath) // 统一前缀
	{
		g.POST("/user", userApi.Create)
		g.DELETE("/user", userApi.Delete)
		g.PUT("/user", userApi.Update)
		g.GET("/user", userApi.Get)
		g.GET("/users", userApi.List)
	}
}
