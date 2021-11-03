package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/apierr"
	"github.com/mittacy/ego-layout/internal/service"
	"github.com/mittacy/ego-layout/internal/transform"
	"github.com/mittacy/ego-layout/internal/validator/userValidator"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/pkg/response"
)

type User struct {
	service   service.User
	transform transform.User
	logger    *log.Logger
}

func NewUser() User {
	l := log.New("user")
	return User{
		logger:    l,
		service:   service.NewUser(l),
		transform: transform.NewUser(l),
	}
}

func (ctl *User) GetUser(c *gin.Context) {
	req := userValidator.GetReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidateErr(c, err)
		return
	}

	user, err := ctl.service.GetUserById(req.Id)
	if err != nil {
		response.CheckErrAndLog(c, ctl.logger, "get user", err, apierr.ErrUserNoExist)
		return
	}

	ctl.transform.GetUserReply(c, user)
}
