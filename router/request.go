package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/pkg/log"
	"time"
)

func InitRequestLog(r *gin.Engine) {
	requestLog := log.New("request")
	r.Use(ginzap.Ginzap(requestLog.GetZap(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(requestLog.GetZap(), true))
}
