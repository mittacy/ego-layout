//+build wireinject

package router

import (
	"github.com/mittacy/ego-layout/app/api"
	"github.com/mittacy/ego-layout/app/data"
	"github.com/mittacy/ego-layout/app/service"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserApi(db *gorm.DB, cache *redis.Pool) api.User {
	wire.Build(data.NewUser, service.NewUser, api.NewUser)
	return api.User{}
}