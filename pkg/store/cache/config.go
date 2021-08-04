package cache

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

const (
	RedisConnPrefix = "redis"		// 配置文件中的前缀
)

var (
	GlobalRedisConf Redis
)

func InitCache() {
	if err := viper.UnmarshalKey(RedisConnPrefix, &GlobalRedisConf); err != nil {
		panic(fmt.Sprintf("cache init err: %s", err))
	}
}

type Redis struct {
	Expire    int64 `mapstructure:"expire"`
	Deviation int64 `mapstructure:"deviation"`
}

type RedisConfig struct {
	Network         string        `mapstructure:"network"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Password        string        `mapstructure:"password"`
	DB              string        `mapstructure:"db"`
	MaxIdle         int           `mapstructure:"maxIdle"`
	MaxActive       int           `mapstructure:"maxActive"`
	IdleTimeout     time.Duration `mapstructure:"idleTimeout"`
	Wait            bool          `mapstructure:"wait"`
	MaxConnLifeTime time.Duration `mapstructure:"maxConnLifeTime"`
}
