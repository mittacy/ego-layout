package service

import (
	"github.com/mittacy/ego-layout/app/model"
	"github.com/mittacy/ego-layout/pkg/log"
)

type User struct {
	userData IUserData
	logger   *log.Logger
}

func NewUser(userData IUserData, logger *log.Logger) User {
	return User{
		userData: userData,
		logger:   logger,
	}
}

type IUserData interface {
	Insert(user *model.User) error
	Delete(id int64) error
	UpdateById(user model.User, updateFields []string) error
	Select(id int64) (*model.User, error)
	List(fields []string, page, pageSize int) ([]model.User, error)
	SelectSum() (int64, error)
}

func (ctl *User) Create(user model.User) (int64, error) {
	if err := ctl.userData.Insert(&user); err != nil {
		return 0, nil
	}

	return user.Id, nil
}

func (ctl *User) Delete(id int64) error {
	return ctl.userData.Delete(id)
}

func (ctl *User) UpdateInfo(user model.User) error {
	fields := []string{"name"}

	return ctl.userData.UpdateById(user, fields)
}

func (ctl *User) UpdatePassword(id int64, password string) error {
	fields := []string{"name"}
	user := model.User{
		Id:       id,
		Password: password, // 加密等操作
	}

	return ctl.userData.UpdateById(user, fields)
}

func (ctl *User) Get(id int64) (*model.User, error) {
	return ctl.userData.Select(id)
}

func (ctl *User) List(page, pageSize int) ([]model.User, int64, error) {
	fields := []string{"id", "name", "created_at", "updated_at"}

	list, err := ctl.userData.List(fields, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalSize, err := ctl.userData.SelectSum()
	if err != nil {
		return nil, 0, err
	}

	return list, totalSize, nil
}
