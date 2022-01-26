package transform

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/mittacy/ego-layout/internal/model"
	"github.com/mittacy/ego-layout/internal/validator/userValidator"
	"github.com/mittacy/ego-layout/pkg/response"
	"github.com/mittacy/log"
)

var User userTransform

type userTransform struct {
	logger *log.Logger
}

func init() {
	l := log.New("user")

	User = userTransform{
		logger: l,
	}
}

// GetUserReply 用户详情响应
// @param data 数据库数据
func (ctl *userTransform) GetUserReply(c *gin.Context, req interface{}, data *model.User) {
	replyUser := userValidator.GetReply{}

	if err := copier.Copy(&replyUser, data); err != nil {
		ctl.logger.ErrorwWithTrace(c, "copier", "req", req, "err", err)
		response.Unknown(c)
		return
	}

	res := map[string]interface{}{
		"user": replyUser,
	}

	response.Success(c, res)
}
