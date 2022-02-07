package config

import (
	"github.com/mittacy/ego/library/eRedis"
	"github.com/spf13/viper"
)

var redisConfigs map[string]eRedis.Conf

func InitRedis() {
	redisConfigs = map[string]eRedis.Conf{
		"localhost": {
			Host:        viper.GetString("REDIS_LOCALHOST_RW_HOST"),
			Password:    viper.GetString("REDIS_LOCALHOST_RW_PASSWORD"),
			Port:        viper.GetInt("REDIS_LOCALHOST_RW_PORT"),
			PoolSize:    viper.GetInt("REDIS_LOCALHOST_POOL_SIZE"),
			MinIdleConn: viper.GetInt("REDIS_LOCALHOST_MIN_IDLE_CONN"),
			IdleTimeout: viper.GetDuration("REDIS_LOCALHOST_IDLE_TIMEOUT"),
		},
	}
	eRedis.InitRedis(redisConfigs)
}
