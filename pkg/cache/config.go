package cache

import "github.com/spf13/viper"

type Conf struct {
	Host     string
	Password string
	Port     int
}

func GetConfig(name string) (Conf, bool) {
	var conf Conf

	switch name {
	case "localhost":
		conf = Conf{
			Host:     viper.GetString("REDIS_LOCALHOST_RW_HOST"),
			Password: viper.GetString("REDIS_LOCALHOST_RW_PASSWORD"),
			Port:     viper.GetInt("REDIS_LOCALHOST_RW_PORT"),
		}
	default:
		return Conf{}, false
	}

	return conf, true
}
