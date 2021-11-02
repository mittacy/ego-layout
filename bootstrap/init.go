package bootstrap

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/pkg/config"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/pkg/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
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

	// 2. 设置GIN运行模式
	appEnv := viper.GetString("APP_ENV")
	gin.SetMode(appEnv)

	// 3. 初始化日志
	logPath := viper.GetString("LOG_PATH")
	logLevel := zapcore.Level(viper.GetInt("LOG_LOW_LEVEL"))
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
		log.WithGlobalFields(globalFields...))

	// 4. 初始化Mysql配置
	mysql.InitConfig()
}
