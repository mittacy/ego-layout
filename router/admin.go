package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/app/adminApi"
)

func InitAdminRouter(r *gin.Engine) {
	// 控制器
	AdminUser := adminApi.NewUser()

	globalPath := "/admin/api"
	g := r.Group(globalPath)
	{
		g.GET("/user/ping", AdminUser.Ping)
	}
}
