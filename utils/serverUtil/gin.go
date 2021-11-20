package serverUtil

import "github.com/gin-gonic/gin"

// AppEnvToGinEnv 转化服务运行环境为GIN的服务运行环境
// @param appEnv development/test/production
// @return string debug/test/release
func AppEnvToGinEnv(appEnv string) string {
	switch appEnv {
	case "development":
		return gin.DebugMode
	case "test":
		return gin.TestMode
	case "production":
		return gin.ReleaseMode
	default:
		return gin.DebugMode
	}
}
