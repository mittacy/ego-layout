package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/interface/api"
)

func InitAdminRouter(r *gin.Engine) {
	// 控制器
	User := api.NewUser()

	globalPath := "/admin/api"
	g := r.Group(globalPath)
	{
		g.GET("/user/info", User.GetUser)
	}
}
