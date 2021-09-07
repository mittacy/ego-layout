package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/pkg/checker"
	"github.com/mittacy/ego-layout/pkg/config"
	"github.com/mittacy/ego-layout/pkg/jwt"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/pkg/store/cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() {
	// 1. 初始化配置文件
	config.InitViper()

	// 2. 设置gin的运行模式
	gin.SetMode(config.ServerConfig.Env)

	// 3. 初始化日志库
	logConf := log.Config{
		ServerName:   viper.GetString("server.name"),
		Path:         viper.GetString("log.path"),
		LowLevel:     log.Level(viper.GetInt("log.lowLevel")),
		LogInConsole: false,
	}
	if gin.Mode() == gin.DebugMode {
		logConf.LogInConsole = true
	} else {
		logConf.LogInConsole = false
	}
	log.Init(logConf)

	// 4. 初始化校验翻译器
	if err := checker.InitTrans(); err != nil {
		log.Panic("初始化校验翻译器失败", zap.String("reason", err.Error()))
	}

	// 5. 初始化 Cache 配置
	cache.Init()

	// 5. 初始化token
	tokenCache := cache.ConnCustomRedis("REDISKEY", "token")
	jwt.InitToken(tokenCache)
}
