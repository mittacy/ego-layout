package cache

import "github.com/spf13/viper"

type Conf struct {
	Host     string
	Password string
	Port     int
}

var (
	cacheConfig map[string]Conf
)

func InitConfig() {
	cacheConfig = map[string]Conf{
		"localhost": {
			Host:     viper.GetString("REDIS_LOCALHOST_RW_HOST"),
			Port:     viper.GetInt("REDIS_LOCALHOST_RW_PORT"),
			Password: viper.GetString("REDIS_LOCALHOST_RW_PASSWORD"),
		},
	}
}
