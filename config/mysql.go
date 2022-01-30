package config

import (
	"github.com/mittacy/ego/library/mysql"
	"github.com/spf13/viper"
)

var mysqlConfigs map[string]mysql.Conf

func InitGorm() {
	mysqlConfigs = map[string]mysql.Conf{
		"localhost": {
			Host:     viper.GetString("DB_CORE_RW_HOST"),
			Port:     viper.GetInt("DB_CORE_RW_PORT"),
			Database: viper.GetString("DB_CORE_RW_DATABASE"),
			User:     viper.GetString("DB_CORE_RW_USERNAME"),
			Password: viper.GetString("DB_CORE_RW_PASSWORD"),
		},
	}
	mysql.InitGorm(mysqlConfigs,
		mysql.WithLogName("mysql"),
		mysql.WithLogSlowThreshold(viper.GetDuration("GORM_SLOW_LOG_THRESHOLD")),
		mysql.WithLogIgnoreRecordNotFound(true))
}
