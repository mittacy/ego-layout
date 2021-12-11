package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

// NewClientByName 直接通过配置名字获取新客户端
// @param name 配置名
// @param db 使用哪个数据库
// @return *redis.Client
func NewClientByName(name string, db int) *redis.Client {
	if conf, ok := GetConfig(name); ok {
		return NewClient(conf, db)
	}

	log.Panicf("%s 配置不存在, 请在 cacheConfig 中配置", name)
	return nil
}

// NewClient 获取新客户端
// @param conf 配置名
// @param db 使用哪个数据库
// @return *redis.Client
func NewClient(conf Conf, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       db,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Panicf("连接数据库失败, 检查配置, err: %s, conf: %+v", err, conf)
	}

	return rdb
}
