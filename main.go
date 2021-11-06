package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/app/task"
	"github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/router"
	"github.com/spf13/viper"
	"net/http"
	"syscall"
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

	// 启动定时任务
	go func() {
		task.StartTasks()
	}()

	if err := graceServe(r); err != nil {
		panic(err)
	}
}

// graceServe 启动服务
// 支持平滑重启，重新编译生成同名可执行文件后，执行 kill -1 pid 即可平滑重启
func graceServe(r *gin.Engine) error {
	endless.DefaultReadTimeOut = time.Second * viper.GetDuration("APP_READ_TIMEOUT")
	endless.DefaultWriteTimeOut = time.Second * viper.GetDuration("APP_WRITE_TIMEOUT")
	endless.DefaultMaxHeaderBytes = 1 << 20
	port := ":" + viper.GetString("APP_PORT")

	server := endless.NewServer(port, r)

	server.BeforeBegin = func(add string) {
		log.Sugar().Infof("Actual pid is %d", syscall.Getpid())
	}

	log.Sugar().Infof("监听端口%s", port)

	return server.ListenAndServe()
}

// serve 启动服务
func serve(r *gin.Engine) error {
	s := &http.Server{
		Addr:           ":" + viper.GetString("APP_PORT"),
		Handler:        r,
		ReadTimeout:    time.Second * viper.GetDuration("APP_READ_TIMEOUT"),
		WriteTimeout:   time.Second * viper.GetDuration("APP_WRITE_TIMEOUT"),
		MaxHeaderBytes: 1 << 20,
	}

	log.Sugar().Infof("监听端口%s", s.Addr)

	return s.ListenAndServe()
}
