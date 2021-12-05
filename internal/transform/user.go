package transform

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/mittacy/ego-layout/internal/model"
	"github.com/mittacy/ego-layout/internal/validator/userValidator"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/pkg/response"
)

type User struct {
	logger *log.Logger
}

func NewUser(logger *log.Logger) User {
	return User{logger: logger}
}

// GetUserReply 用户详情响应
// @param data 数据库数据
func (ctl *User) GetUserReply(c *gin.Context, req interface{}, data *model.User) {
	replyUser := userValidator.GetReply{}

	if err := copier.Copy(&replyUser, data); err != nil {
		ctl.logger.CopierErrLog(err, req)
		response.Unknown(c)
		return
	}

	res := map[string]interface{}{
		"user": replyUser,
	}

	response.Success(c, res)
}
