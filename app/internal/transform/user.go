package transform

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/mittacy/ego-layout/app/internal/model"
	"github.com/mittacy/ego-layout/app/internal/validator/userValidator"
	"github.com/mittacy/ego/library/gin/response"
	"github.com/mittacy/ego/library/log"
)

var User userTransform

type userTransform struct{}

// GetUserReply 用户详情响应
// @param data 数据库数据
func (ctl *userTransform) GetUserReply(c *gin.Context, req interface{}, data *model.User) {
	replyUser := userValidator.GetReply{}

	if err := copier.Copy(&replyUser, data); err != nil {
		log.New("user").ErrorwWithTrace(c, "copier", "req", req, "err", err)
		response.Unknown(c)
		return
	}

	res := map[string]interface{}{
		"user": replyUser,
	}

	response.Success(c, res)
}
