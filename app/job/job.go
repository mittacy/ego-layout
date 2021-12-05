package job

import (
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego-layout/internal/job/exampleJob"
	"github.com/mittacy/ego-layout/pkg/async"
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

// Serve 异步任务服务
func Serve(stop <-chan struct{}) error {
	jobs := Jobs()

	srv := asynq.NewServer(
		async.GetDefaultRedisConnOpt(),
		async.GetDefaultServerConfig(),
	)

	go func() {
		<-stop
		srv.Shutdown()
	}()

	mux := asynq.NewServeMux()
	for _, v := range jobs {
		mux.Handle(v.TypeName, v.Handler)
	}

	return srv.Run(mux)
}
