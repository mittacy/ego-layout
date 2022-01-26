package bootstrap

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/middleware"
	"github.com/mittacy/ego-layout/pkg/config"
	"github.com/mittacy/ego-layout/utils/serverUtil"
	"github.com/mittacy/log"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func init() {
	/* 环境变量解析
	 	-config 配置文件路径 example: -config ./.env.development
		-port 服务监听端口 example: -port 10244
		-env 服务监听端口 example: -env development/test/production
	*/
	configPath := flag.String("config", ".env.development", "配置文件名")
	serverEnv := flag.String("env", "", "服务环境")
	serverPort := flag.String("port", "10244", "服务监听端口")
	flag.Parse()

	// 1. 初始化配置文件
	config.Init(*configPath)
	// 命令行参数覆盖env配置
	viper.Set("APP_PORT", *serverPort)
	if *serverEnv != "" {
		viper.Set("APP_ENV", *serverEnv)
	}

	// 2. 设置GIN运行模式
	gin.SetMode(serverUtil.AppEnvToGinEnv(viper.GetString("APP_ENV")))

	// 3. 初始化日志
	logPath := viper.GetString("LOG_PATH")
	logLevel := zapcore.Level(viper.GetInt("LOG_LOW_LEVEL"))
	logEncoderJson := viper.GetBool("LOG_ENCODER_JSON")
	logInConsole := false
	if gin.Mode() == gin.DebugMode {
		logInConsole = true
	}
	globalFields := []zapcore.Field{
		{
			Key:    "module_name",
			Type:   zapcore.StringType,
			String: viper.GetString("APP_NAME"),
		},
	}

	log.SetDefaultConf(
		log.WithPath(logPath),
		log.WithTimeFormat("2006-01-02 15:04:05"),
		log.WithLevel(logLevel),
		log.WithPreName("biz_"),
		log.WithEncoderJSON(logEncoderJson),
		log.WithFields(globalFields...),
		log.WithLogInConsole(logInConsole),
		log.WithRequestIdKey(middleware.RequestIdKey))
}
