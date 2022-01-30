package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/app/api"
)

func InitAdminRouter(r *gin.Engine) {
	globalPath := "/admin/api"
	g := r.Group(globalPath)
	{
		g.GET("/user/info", api.User.GetUser)
	}
}
