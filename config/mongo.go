package config

import (
	"github.com/mittacy/ego/library/eMongo"
	"github.com/spf13/viper"
)

var mongoConfigs map[string]eMongo.Conf

func InitMongo() {
	mongoConfigs = map[string]eMongo.Conf{
		"localhost": {
			Host:     viper.GetString("MONGO_LOCALHOST_RW_HOST"),
			Port:     viper.GetInt("MONGO_LOCALHOST_RW_PORT"),
			Database: viper.GetString("MONGO_LOCALHOST_RW_DATABASE"),
			User:     viper.GetString("MONGO_LOCALHOST_RW_USERNAME"),
			Password: viper.GetString("MONGO_LOCALHOST_RW_PASSWORD"),
		},
	}

	eMongo.Init(mongoConfigs)
}
