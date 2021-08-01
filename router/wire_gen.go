// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package router

import (
	"github.com/mittacy/ego-layout/app/api"
	"github.com/mittacy/ego-layout/app/data"
	"github.com/mittacy/ego-layout/app/service"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitUserApi(db *gorm.DB, cache *redis.Pool) api.User {
	iUserData := data.NewUser(db, cache)
	iUserService := service.NewUser(iUserData)
	user := api.NewUser(iUserService)
	return user
}
