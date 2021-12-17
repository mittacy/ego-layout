package service

import (
	"github.com/mittacy/ego-layout/internal/data"
	"github.com/mittacy/ego-layout/internal/model"
	"github.com/mittacy/ego-layout/pkg/log"
)

// 一般情况下service应该只包含并调用自己的data模型，需要其他服务的功能请service.Xxx调用服务而不是引入其他data模型
// User 用户服务
var User userService

func init() {
	l := log.New("user")

	User = userService{
		logger: l,
		data: data.NewUser(l),
	}
}

type userService struct {
	logger *log.Logger
	data   data.User
}

func (ctl *userService) GetById(id int64) (*model.User, error) {
	return ctl.data.GetById(id)
}
