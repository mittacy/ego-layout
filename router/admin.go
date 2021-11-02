package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/app/adminApi"
)

func InitAdminRouter(r *gin.Engine) {
	globalPath := "/admin/api"

	g := r.Group(globalPath)
	{
		g.GET("/user/ping", adminApi.User.Ping)
	}
}
