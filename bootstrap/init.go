package bootstrap

import (
	"flag"
	"github.com/mittacy/ego-layout/pkg/config"
	"github.com/spf13/viper"
)

func Init() {
	/* 环境变量解析
	 	-config 配置文件路径 example: ./.env.development
		-port 服务监听端口 example: 10244
		-env 服务监听端口 example: debug/test/release
	 */
	configPath := flag.String("config", "./.env.development", "配置文件名")
	serverPort := flag.String("port", "10244", "服务监听端口")
	serverEnv := flag.String("env", "debug", "服务环境")
	flag.Parse()

	// 1. 初始化配置文件
	config.Init(*configPath)
	viper.Set("APP_PORT", *serverPort)
	viper.Set("APP_ENV", *serverEnv)
}
