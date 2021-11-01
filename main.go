package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/bootstrap"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func init() {
	bootstrap.Init()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "", "msg": "success"})
	})
	s := &http.Server{
		Addr: ":" + viper.GetString("APP_PORT"),
		Handler: r,
		ReadTimeout: time.Second * viper.GetDuration("APP_READ_TIMEOUT"),
		WriteTimeout: time.Second * viper.GetDuration("APP_WRITE_TIMEOUT"),
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
