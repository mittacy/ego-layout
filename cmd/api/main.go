package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	_ "github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/interface/task"
	"github.com/mittacy/ego-layout/middleware"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/router"
	"github.com/spf13/viper"
	"net/http"
	"runtime"
	"strings"
	"syscall"
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

	// Windows不支持信号量无法使用endless自动重启，使用普通服务
	var server func(r *gin.Engine) error
	if runtime.GOOS == "windows" {
		server = serve
	} else {
		server = graceServe
	}

	// 启动API服务
	if err := server(r); err != nil {
		if strings.Contains(err.Error(), "use of closed network connection") {
			log.Infof("执行了kill端口")
		} else {
			log.Panicf("API服务异常退出: %+v", err)
		}
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
