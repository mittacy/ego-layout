package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/apierr"
	"github.com/mittacy/ego-layout/internal/service"
	"github.com/mittacy/ego-layout/internal/transform"
	"github.com/mittacy/ego-layout/internal/validator/userValidator"
	"github.com/mittacy/ego/library/gin/response"
	"github.com/mittacy/ego/library/log"
)

var User userApi

type userApi struct{}

func (ctl *userApi) Ping(c *gin.Context) {
	response.Success(c, "success")
}

func (ctl *userApi) GetUser(c *gin.Context) {
	req := userValidator.GetReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidateErr(c, err)
		return
	}

	user, err := service.User.GetById(c, req.Id)
	if err != nil {
		response.CheckErrAndLog(c, log.New("user"), req, "getUser", err, apierr.UserNoExist)
		return
	}

	transform.User.GetUserReply(c, req, user)
}
