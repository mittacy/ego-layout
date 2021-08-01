package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/pkg/config"
	"github.com/mittacy/ego-layout/router"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func init() {
	bootstrap.Init()
}

func main() {
	r := gin.New()

	// 初始化路由
	router.InitRouter(r)

	serverConfig := config.Global.Server
	s := &http.Server{
		Addr: ":" + strconv.Itoa(serverConfig.Port),
		Handler: r,
		ReadTimeout: time.Second * serverConfig.ReadTimeout,
		WriteTimeout: time.Second * serverConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	zap.S().Infof("监听端口:%d", serverConfig.Port)

	s.ListenAndServe()
}
