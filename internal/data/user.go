package data

import "github.com/mittacy/ego-layout/internal/model"

type User struct{}

func NewUser() User {
	return User{}
}

func (ctl *User) GetById(id int64) (*model.User, error) {
	return &model.User{
		Id:        id,
		Name:      "Mittacy Chen",
		Info:      "Yoyo",
		Password:  "123456",
		Deleted:   0,
		CreatedAt: 1635823685,
		UpdatedAt: 1635823685,
	}, nil
}
