package main

import (
	"github.com/hibiken/asynq"
	_ "github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/interface/job/exampleJob"
	"github.com/mittacy/ego-layout/pkg/async"
	"github.com/mittacy/ego-layout/pkg/log"
	"strings"
)

type Job struct {
	TypeName string
	Handler  asynq.Handler
}

// Jobs 异步任务列表，新增的任务在此添加即可
func Jobs() []Job {
	return []Job{
		{exampleJob.TypeName, exampleJob.NewProcessor()},
	}
}

func main() {
	jobs := Jobs()

	srv := asynq.NewServer(
		async.GetDefaultRedisConnOpt(),
		async.GetDefaultServerConfig(),
	)

	mux := asynq.NewServeMux()
	for _, v := range jobs {
		mux.Handle(v.TypeName, v.Handler)
	}

	if err := srv.Run(mux); err != nil {
		if strings.Contains(err.Error(), "use of closed network connection") {
			log.Infof("执行了kill端口")
		} else {
			log.Panicf("队列服务异常退出: %+v", err)
		}
	}
}
