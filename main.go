package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/app/api"
	"github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/router"
)

func init() {
	bootstrap.Init()
}

func main() {
	done := make(chan error, 1) // 启动监听服务数
	stop := make(chan struct{})

	// 启动异步任务服务
	//go func() {
	//	done <- job.Serve(stop)
	//}()

	// 启动API服务
	go func() {
		r := gin.New()
		r.Use(gin.Recovery())

		router.InitRequestLog(r)
		router.InitRouter(r)
		router.InitAdminRouter(r)

		//task.StartTasks()	// 启动定时任务

		done <- api.GraceServe(r, stop)
	}()

	// 监听多个服务，一个退出则全部执行安全退出
	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			log.Errorf("服务异常退出, err: %s", err)
		}

		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
