package adminApi

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/internal/service"
	"github.com/mittacy/ego-layout/pkg/log"
	"net/http"
)

var User = NewUser()

type UserApi struct {
	logger  *log.Logger
	service service.User
}

func NewUser() UserApi {
	return UserApi{
		logger:  log.New("admin-user"),
		service: service.NewUser(),
	}
}

func (ctl *UserApi) Ping(c *gin.Context) {
	resData := "admin user api ping success"
	c.JSON(http.StatusOK, gin.H{
		"data": resData,
		"msg":  "success",
	})
}
