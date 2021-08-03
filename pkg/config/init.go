package config

import (
	"flag"
	"github.com/spf13/viper"
)

var Global global

type global struct {
	Server `mapstructure:"server"`
	Redis  `mapstructure:"redis"`
	Jwt    `mapstructure:"jwt"`
}

// InitViper 初始化Viper
func InitViper() {
	// 从命令行读取配置文件路径
	path := flag.String("config", "./default.yaml", "配置文件名")
	flag.Parse()

	// 全局初始化 Viper
	viper.SetConfigFile(*path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 解析到全局配置文件中
	if err := viper.Unmarshal(&Global); err != nil {
		panic(err)
	}
}
