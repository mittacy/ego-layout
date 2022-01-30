package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/config"
	"github.com/mittacy/ego-layout/utils/serverUtil"
	"github.com/spf13/viper"
)

// InitHTTP http配置初始化
// @param confPath
// @param env
// @param port
func InitHTTP(confPath, env string, port int) {
	// conf
	InitViper(confPath, env, port)

	// gin run env
	gin.SetMode(serverUtil.AppEnvToGinEnv(viper.GetString("APP_ENV")))

	// log
	InitLog()

	// configs
	config.InitGorm()
	config.InitRedis()
}

