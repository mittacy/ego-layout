package service

import (
	"github.com/mittacy/ego-layout/app/api"
	"github.com/mittacy/ego-layout/app/model"
)

type User struct {
	userService IUserService
}

// 实现api层中的各个service接口的构建方法

func NewUser(userService IUserService) api.IUserService {
	return &User{
		userService: userService,
	}
}

type IUserService interface {
	Insert(user *model.User) error
	Delete(id int64) error
	UpdateById(user model.User, updateFields []string) error
	Select(id int64) (*model.User, error)
	List(fields []string, page, pageSize int) ([]model.User, error)
	SelectSum() (int64, error)
}

func (ctl *User) Create(user model.User) (int64, error) {
	if err := ctl.userService.Insert(&user); err != nil {
		return 0, nil
	}

	return user.Id, nil
}

func (ctl *User) Delete(id int64) error {
	return ctl.userService.Delete(id)
}

func (ctl *User) UpdateInfo(user model.User) error {
	fields := []string{"name"}

	return ctl.userService.UpdateById(user, fields)
}

func (ctl *User) UpdatePassword(id int64, password string) error {
	fields := []string{"name"}
	user := model.User{
		Id:       id,
		Password: password, // 加密等操作
	}

	return ctl.userService.UpdateById(user, fields)
}

func (ctl *User) Get(id int64) (*model.User, error) {
	return ctl.userService.Select(id)
}

func (ctl *User) List(page, pageSize int) ([]model.User, int64, error) {
	fields := []string{"id", "name", "created_at", "updated_at"}

	list, err := ctl.userService.List(fields, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalSize, err := ctl.userService.SelectSum()
	if err != nil {
		return nil, 0, err
	}

	return list, totalSize, nil
}
