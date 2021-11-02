package data

import (
	"github.com/mittacy/ego-layout/apierr"
	"github.com/mittacy/ego-layout/internal/model"
	"github.com/mittacy/ego-layout/pkg/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser() User {
	return User{
		db: mysql.NewClientByName("localhost"),
	}
}

func (ctl *User) GetById(id int64) (*model.User, error) {
	user := model.User{}
	if err := ctl.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apierr.ErrUserNoExist
		}

		return nil, errors.WithStack(err)
	}

	return &user, nil
}
