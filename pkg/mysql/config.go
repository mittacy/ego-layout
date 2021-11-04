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

var (
	dbConfig map[string]Conf
)

func InitConfig() {
	dbConfig = map[string]Conf{
		"localhost": {
			Host:     viper.GetString("DB_CORE_RW_HOST"),
			Port:     viper.GetInt("DB_CORE_RW_PORT"),
			Database: viper.GetString("DB_DATABASE_RESOURCES"),
			User:     viper.GetString("DB_CORE_RW_USERNAME"),
			Password: viper.GetString("DB_CORE_RW_PASSWORD"),
		},
	}
}
