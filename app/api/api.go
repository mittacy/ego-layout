package api

import (
	"context"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/spf13/viper"
	"net/http"
	"syscall"
	"time"
)

// GraceServe 启动服务
// 支持平滑重启，重新编译生成同名可执行文件后，执行 kill -1 pid 即可平滑重启
func GraceServe(r *gin.Engine, stop <-chan struct{}) error {
	endless.DefaultReadTimeOut = time.Second * viper.GetDuration("APP_READ_TIMEOUT")
	endless.DefaultWriteTimeOut = time.Second * viper.GetDuration("APP_WRITE_TIMEOUT")
	endless.DefaultMaxHeaderBytes = 1 << 20
	port := ":" + viper.GetString("APP_PORT")

	server := endless.NewServer(port, r)

	server.BeforeBegin = func(add string) {
		log.Sugar().Infof("Actual pid is %d", syscall.Getpid())
	}

	log.Sugar().Infof("监听端口%s", port)

	go func() {
		<-stop
		server.Shutdown(context.Background())
	}()

	return server.ListenAndServe()
}

// serve 启动服务
func serve(r *gin.Engine, stop <-chan struct{}) error {
	s := &http.Server{
		Addr:           ":" + viper.GetString("APP_PORT"),
		Handler:        r,
		ReadTimeout:    time.Second * viper.GetDuration("APP_READ_TIMEOUT"),
		WriteTimeout:   time.Second * viper.GetDuration("APP_WRITE_TIMEOUT"),
		MaxHeaderBytes: 1 << 20,
	}

	log.Sugar().Infof("监听端口%s", s.Addr)

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}
