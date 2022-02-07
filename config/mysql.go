package config

import (
	"github.com/mittacy/ego/library/eMysql"
	"github.com/spf13/viper"
)

var mysqlConfigs map[string]eMysql.Conf

func InitGorm() {
	mysqlConfigs = map[string]eMysql.Conf{
		"localhost": {
			Host:     viper.GetString("DB_LOCALHOST_RW_HOST"),
			Port:     viper.GetInt("DB_LOCALHOST_RW_PORT"),
			Database: viper.GetString("DB_LOCALHOST_RW_DATABASE"),
			User:     viper.GetString("DB_LOCALHOST_RW_USERNAME"),
			Password: viper.GetString("DB_LOCALHOST_RW_PASSWORD"),
		},
	}
	eMysql.InitGorm(mysqlConfigs,
		eMysql.WithLogName("mysql"),
		eMysql.WithLogSlowThreshold(viper.GetDuration("GORM_SLOW_LOG_THRESHOLD")),
		eMysql.WithLogIgnoreRecordNotFound(true))
}
