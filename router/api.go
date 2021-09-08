package router

import (
	"github.com/mittacy/ego-layout/app/api"
	"github.com/mittacy/ego-layout/app/data"
	"github.com/mittacy/ego-layout/app/service"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/pkg/store/cache"
	"github.com/mittacy/ego-layout/pkg/store/db"
	"gorm.io/gorm"
)

var (
	userApi api.User
)

func InitApi() {
	userApi = InitUserApi(db.ConnectGorm("MYSQLKEY"), cache.NewRedis("REDISKEY", "user"))
}

func InitUserApi(db *gorm.DB, redis *cache.Redis) api.User {
	customLogger := log.New("user")
	userData := data.NewUser(db, redis, customLogger)
	userService := service.NewUser(userData, customLogger)
	return api.NewUser(userService, customLogger)
}
