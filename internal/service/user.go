package service

import (
	"github.com/mittacy/ego-layout/internal/data"
	"github.com/mittacy/ego-layout/internal/model"
)

type User struct {
	data data.User
}

func NewUser() User {
	return User{
		data: data.NewUser(),
	}
}

func (ctl *User) GetUserById(id int64) (*model.User, error) {
	return ctl.data.GetById(id)
}
