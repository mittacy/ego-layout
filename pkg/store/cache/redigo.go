package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

var (
	cachePool map[string]*redis.Pool
)

func init() {
	cachePool = make(map[string]*redis.Pool, 0)
}

// GetRedisPool 获取 redis 连接池
// @param name redis配置名
// @return redis.Conn
func GetRedisPool(name string) *redis.Pool {
	key := fmt.Sprintf("%s.%s", RedisConnPrefix, name)

	if conn, ok := cachePool[key]; ok {
		return conn
	}

	var conf RedisConfig
	if err := viper.UnmarshalKey(key, &conf); err != nil {
		zap.S().Panicf("连接redis失败, 检查%s的配置, err: %s", key, err)
	}

	pool, err := connectRedigoPool(conf)
	if err != nil {
		zap.S().Panicf("连接redis失败, 检查%s的配置, err: %s", key, err)
	}

	cachePool[key] = pool
	return pool
}

func connectRedigoPool(conf RedisConfig) (*redis.Pool, error) {
	pool := redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(conf.Network, fmt.Sprintf("%s:%d", conf.Host, conf.Port))
			if err != nil {
				return nil, err
			}

			if conf.Password != "" {
				if _, err := c.Do("AUTH", conf.Password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", conf.DB); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
		MaxIdle:         conf.MaxIdle,                       // 最大空闲连接数
		MaxActive:       conf.MaxActive,                     // 连接池最大数目,为0则不限制
		IdleTimeout:     conf.IdleTimeout * time.Second,     // 空闲连接超时时间，超过时间的空闲连接会被关闭,为0将不会被关闭,应该设置一个比redis服务端超时时间更短的时间
		Wait:            conf.Wait,                          // 如果为true且已经达到MaxActive的限制，则等待连接池
		MaxConnLifetime: conf.MaxConnLifeTime * time.Second, // 超过该时间关闭连接，如果为0则不根据时间来关闭连接
	}

	return &pool, nil
}
