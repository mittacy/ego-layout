package bootstrap

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/pkg/cache"
	"github.com/mittacy/ego-layout/pkg/config"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/pkg/mysql"
	"github.com/mittacy/ego-layout/utils/serverUtil"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func Init() {
	/* 环境变量解析
	 	-config 配置文件路径 example: ./.env.develop
		-port 服务监听端口 example: 10244
		-env 服务监听端口 example: develop/test/production
	*/
	configPath := flag.String("config", ".env.develop", "配置文件名")
	serverEnv := flag.String("env", "", "服务环境")
	serverPort := flag.String("port", "10244", "服务监听端口")
	flag.Parse()

	// 1. 初始化配置文件
	config.Init(*configPath)
	// 命令行参数覆盖env配置
	viper.Set("APP_PORT", *serverPort)
	if *serverEnv != "" {
		viper.GetString("APP_ENV")
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
			String: "serverName",
		},
	}

	log.SetGlobalConfig(
		log.WithPath(logPath),
		log.WithLevel(logLevel),
		log.WithLogInConsole(logInConsole),
		log.WithGlobalFields(globalFields...),
		log.WithGlobalEncoderJSON(logEncoderJson))

	// 4. 初始化Mysql配置
	mysql.InitConfig()

	// 5. 初始化缓存配置
	cache.InitConfig()
}
