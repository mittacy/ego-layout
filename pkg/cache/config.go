package cache

import (
	"github.com/spf13/viper"
	"time"
)

type Conf struct {
	Host        string
	Password    string
	Port        int
	PoolSize    int
	MinIdleConn int
	IdleTimeout time.Duration
}

func GetConfig(name string) (Conf, bool) {
	var conf Conf

	switch name {
	case "localhost":
		conf = Conf{
			Host:        viper.GetString("REDIS_LOCALHOST_RW_HOST"),
			Password:    viper.GetString("REDIS_LOCALHOST_RW_PASSWORD"),
			Port:        viper.GetInt("REDIS_LOCALHOST_RW_PORT"),
			PoolSize:    viper.GetInt("REDIS_LOCALHOST_POOL_SIZE"),
			MinIdleConn: viper.GetInt("REDIS_LOCALHOST_MIN_IDLE_CONN"),
			IdleTimeout: viper.GetDuration("REDIS_LOCALHOST_IDLE_TIMEOUT"),
		}
	default:
		return Conf{}, false
	}

	return conf, true
}
