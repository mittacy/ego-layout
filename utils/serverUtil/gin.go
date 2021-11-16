package serverUtil

import "github.com/gin-gonic/gin"

// AppEnvToGinEnv 转化服务运行环境为GIN的服务运行环境
// @param appEnv develop/test/production
// @return string debug/test/release
func AppEnvToGinEnv(appEnv string) string {
	switch appEnv {
	case "develop":
		return gin.DebugMode
	case "test":
		return gin.TestMode
	case "production":
		return gin.ReleaseMode
	default:
		return gin.DebugMode
	}
}
