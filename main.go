package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/pkg/config"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/router"
	"net/http"
	"strconv"
	"time"
)

func init() {
	bootstrap.Init()
}

func main() {
	defer log.Sync()

	r := gin.New()

	// 初始化路由
	router.InitRouter(r)

	serverConfig := config.ServerConfig
	s := &http.Server{
		Addr: ":" + strconv.Itoa(serverConfig.Port),
		Handler: r,
		ReadTimeout: time.Second * serverConfig.ReadTimeout,
		WriteTimeout: time.Second * serverConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Infof("监听端口:%d", serverConfig.Port)

	s.ListenAndServe()
}
