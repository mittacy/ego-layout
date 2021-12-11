package service

import (
	"github.com/mittacy/ego-layout/internal/data"
	"github.com/mittacy/ego-layout/internal/model"
	"github.com/mittacy/ego-layout/pkg/log"
)

var User userService

func init() {
	l := log.New("user")

	User = userService{
		logger: l,
	}
}

type userService struct {
	logger *log.Logger
}

func (ctl *userService) GetById(id int64) (*model.User, error) {
	return data.User.GetById(id)
}
