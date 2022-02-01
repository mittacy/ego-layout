package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/app/internal/data"
	"github.com/mittacy/ego-layout/app/internal/model"
)

// 一般情况下service应该只包含并调用自己的data模型，需要其他服务的功能请service.Xxx调用服务而不是引入其他data模型
// User 用户服务
var User = userService{
	data: data.NewUser(),
}

type userService struct {
	data data.User
}

func (ctl *userService) GetById(c *gin.Context, id int64) (*model.User, error) {
	return ctl.data.GetById(c, id)
}
