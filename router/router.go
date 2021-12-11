package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/interface/api"
)

func InitRouter(r *gin.Engine) {
	// 控制器
	User := api.NewUser()

	globalPath := "/api"
	g := r.Group(globalPath)
	{
		g.GET("/user/ping", User.Ping)
	}
}
