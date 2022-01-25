package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/interface/task"
	"github.com/mittacy/ego-layout/middleware"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/router"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestTrace())

	router.InitRequestLog(r)
	router.InitRouter(r)
	router.InitAdminRouter(r)

	// 启动定时任务
	go func() {
		task.StartTasks()
	}()

	// 启动API服务
	if err := serve(r); err != nil {
		if strings.Contains(err.Error(), "use of closed network connection") {
			log.Info("执行了kill端口")
		} else {
			log.Panicw("API服务异常退出", "err", err)
		}
	}
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

	log.Infow("服务启动中", "端口号", s.Addr)

	return s.ListenAndServe()
}
