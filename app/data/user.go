package data

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mittacy/ego-layout/app/model"
	"github.com/mittacy/ego-layout/app/service"
	"github.com/mittacy/ego-layout/pkg/logger"
	"github.com/mittacy/ego-layout/pkg/store/cache"
	"gorm.io/gorm"
)

// 实现service层中的data接口

type User struct {
	db     *gorm.DB
	cache  cache.CustomRedis
	logger *logger.CustomLogger
}

func NewUser(db *gorm.DB, cacheConn *redis.Pool, logger *logger.CustomLogger) service.IUserData {
	c := cache.ConnRedisByPool(cacheConn, "user")

	return &User{
		db:     db,
		cache:  c,
		logger: logger,
	}
}

func (ctl *User) Insert(user *model.User) error {
	// 操作缓存和数据库

	user.Id = 520
	return nil
}

func (ctl *User) Delete(id int64) error {
	// 操作缓存和数据库

	return nil
}

func (ctl *User) UpdateById(user model.User, updateFields []string) error {
	// 操作缓存和数据库

	return nil
}

func (ctl *User) Select(id int64) (*model.User, error) {
	// 操作缓存和数据库

	return &model.User{
		Id:        id,
		Name:      "name",
		Info:      "this is info",
		Password:  "5HfTQ8jeGN5o",
		CreatedAt: 1627455919,
		UpdatedAt: 1627659959,
	}, nil
}

func (ctl *User) List(fields []string, page, pageSize int) ([]model.User, error) {
	// 操作缓存和数据库

	users := []model.User{
		{Id: 1, Name: "Xiao Ming", CreatedAt: 1627455919, UpdatedAt: 1627659959},
		{Id: 2, Name: "Xiao Hong", CreatedAt: 1627455012, UpdatedAt: 1627651203},
	}

	return users, nil
}

func (ctl *User) SelectSum() (int64, error) {
	// 操作缓存和数据库

	return 2, nil
}
