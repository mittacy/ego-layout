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
	iUserService := data.NewUser(db, cache, customLogger)
	apiIUserService := service.NewUser(iUserService, customLogger)
	user := api.NewUser(apiIUserService, customLogger)
	return user
}
