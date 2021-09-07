package router

import (
	"github.com/gomodule/redigo/redis"
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
	userApi = InitUserApi(db.ConnectGorm("MYSQLKEY"), cache.ConnRedis("REDISKEY"))
}

func InitUserApi(db *gorm.DB, cache *redis.Pool) api.User {
	customLogger := log.New("user")
	userData := data.NewUser(db, cache, customLogger)
	userService := service.NewUser(userData, customLogger)
	userApi := api.NewUser(userService, customLogger)
	return userApi
}
