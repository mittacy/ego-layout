package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/middleware"
	"github.com/mittacy/ego/library/log"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func InitLog() {
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
