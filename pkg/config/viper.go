package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func Init(file string) {
	viper.SetConfigType("env")
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("load config file fail: %s", err)
	}

	// 监听配置实时更新
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(e.Name, e.Op, e.String())
		viper.SetConfigFile(e.Name)
		log.Printf("some configuration item in the %s file has changed", e.Name)
	})
	viper.WatchConfig()
}
