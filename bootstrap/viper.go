package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func InitViper(file string, env string, port int) {
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

	// 环境变量覆盖值
	if port != 0 {
		viper.Set("APP_PORT", port)
	}

	if env != "" {
		viper.Set("APP_ENV", env)
	}
}
