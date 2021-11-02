package adminApi

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/internal/service"
	"github.com/mittacy/ego-layout/pkg/log"
	"net/http"
)

type UserApi struct {
	logger  *log.Logger
	service service.User
}

func NewUser() UserApi {
	l := log.New("admin-user")

	return UserApi{
		logger:  l,
		service: service.NewUser(l),
	}
}

func (ctl *UserApi) Ping(c *gin.Context) {
	resData := "admin user api ping success"
	c.JSON(http.StatusOK, gin.H{
		"data": resData,
		"msg":  "success",
	})
}
