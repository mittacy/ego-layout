package service

import (
	"github.com/mittacy/ego-layout/internal/data"
	"github.com/mittacy/ego-layout/internal/model"
	"github.com/mittacy/ego-layout/pkg/log"
)

type User struct {
	data   data.User
	logger *log.Logger
}

func NewUser(logger *log.Logger) User {
	return User{
		data:   data.NewUser(logger),
		logger: logger,
	}
}

func (ctl *User) GetUserById(id int64) (*model.User, error) {
	return ctl.data.GetById(id)
}
