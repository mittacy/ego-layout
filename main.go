package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/router"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func init() {
	bootstrap.Init()
}

func main() {
	r := gin.New()
	router.InitRequestLog(r)
	router.InitRouter(r)
	router.InitAdminRouter(r)

	s := &http.Server{
		Addr:           ":" + viper.GetString("APP_PORT"),
		Handler:        r,
		ReadTimeout:    time.Second * viper.GetDuration("APP_READ_TIMEOUT"),
		WriteTimeout:   time.Second * viper.GetDuration("APP_WRITE_TIMEOUT"),
		MaxHeaderBytes: 1 << 20,
	}

	log.Sugar().Infof("监听端口%s", s.Addr)

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
