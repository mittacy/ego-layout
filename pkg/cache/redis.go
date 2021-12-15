package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var (
	cachePool map[string]*redis.Client
)

func init() {
	cachePool = make(map[string]*redis.Client, 0)
}

// NewClientByName 直接通过配置名字获取新客户端
// @param name 配置名
// @param db 使用哪个数据库
// @return *redis.Client
func NewClientByName(name string, db int) *redis.Client {
	cacheName := cachePoolName(name, db)
	if c, ok := cachePool[cacheName]; ok {
		return c
	}

	if conf, ok := GetConfig(name); ok {
		client := NewClient(conf, db)
		cachePool[cacheName] = client
		return client
	}

	log.Panicf("%s 配置不存在, 请在 pkg/cache/config.go GetConfig() 中配置", name)
	return nil
}

// NewClient 获取新客户端
// @param conf 配置名
// @param db 使用哪个数据库
// @return *redis.Client
func NewClient(conf Conf, db int) *redis.Client {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       db,
	}

	if conf.PoolSize > 0 { // 最大连接数
		options.PoolSize = conf.PoolSize
	}
	if conf.MinIdleConn > 0 { // 最小空闲连接数
		options.MinIdleConns = conf.MinIdleConn
	}
	if conf.IdleTimeout > 0 { // 空闲时间(秒)
		options.IdleTimeout = conf.IdleTimeout * time.Second
	}

	rdb := redis.NewClient(options)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Panicf("连接数据库失败, 检查配置, err: %s, conf: %+v", err, conf)
	}

	return rdb
}

func cachePoolName(name string, db int) string {
	return fmt.Sprintf("%s:%d", name, db)
}
