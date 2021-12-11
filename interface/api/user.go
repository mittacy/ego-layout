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

var User userApi

func init() {
	l := log.New("user")

	User = userApi{
		logger: l,
	}
}

type userApi struct {
	logger *log.Logger
}

func (ctl *userApi) Ping(c *gin.Context) {
	response.Success(c, "success")
}

func (ctl *userApi) GetUser(c *gin.Context) {
	req := userValidator.GetReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidateErr(c, err)
		return
	}

	user, err := service.User.GetById(req.Id)
	if err != nil {
		response.CheckErrAndLog(c, ctl.logger, req, "get user", err, apierr.ErrUserNoExist)
		return
	}

	transform.User.GetUserReply(c, req, user)
}
