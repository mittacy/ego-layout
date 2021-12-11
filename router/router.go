package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/interface/api"
)

func InitRouter(r *gin.Engine) {
	globalPath := "/api"
	g := r.Group(globalPath)
	{
		g.GET("/user/ping", api.User.Ping)
	}
}
