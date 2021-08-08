package router

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mittacy/ego-layout/app/api"
	"github.com/mittacy/ego-layout/app/data"
	"github.com/mittacy/ego-layout/app/service"
	"github.com/mittacy/ego-layout/pkg/logger"
	"gorm.io/gorm"
)

func InitUserApi(db *gorm.DB, cache *redis.Pool) api.User {
	customLogger := logger.NewCustomLogger("user")
	userData := data.NewUser(db, cache, customLogger)
	userService := service.NewUser(userData, customLogger)
	userApi := api.NewUser(userService, customLogger)
	return userApi
}
