package mysql

import "github.com/spf13/viper"

type Conf struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
	Params   string
}

func GetConfig(name string) (Conf, bool) {
	var conf Conf

	switch name {
	case "localhost":
		conf = Conf{
			Host:     viper.GetString("DB_CORE_RW_HOST"),
			Port:     viper.GetInt("DB_CORE_RW_PORT"),
			Database: viper.GetString("DB_DATABASE_RESOURCES"),
			User:     viper.GetString("DB_CORE_RW_USERNAME"),
			Password: viper.GetString("DB_CORE_RW_PASSWORD"),
		}
	default:
		return Conf{}, false
	}

	return conf, true
}
