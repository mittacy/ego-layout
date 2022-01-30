package job

import (
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego-layout/app/job/exampleJob/exampleJobProcessor"
	"github.com/mittacy/ego-layout/app/job/exampleJob/exampleJobTask"
	"github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/pkg/async"
	"github.com/mittacy/ego/library/log"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "job",
	Short: "start the async job server",
	Long:  "start the async job server. Example: server start job",
	Run:   run,
}

var (
	conf string
	env  string
	l    *log.Logger
)

func init() {
	Cmd.Flags().StringVarP(&conf, "conf", "c", ".env.development", "配置文件路径")
	Cmd.Flags().StringVarP(&env, "env", "e", "development", "运行环境")
}

type Job struct {
	TypeName string
	Handler  asynq.Handler
}

// Jobs 异步任务列表，新增的任务在此添加即可
func Jobs() []Job {
	return []Job{
		{exampleJobTask.TypeName, exampleJobProcessor.NewProcessor()},
	}
}

func run(cmd *cobra.Command, args []string) {
	bootstrap.InitJob(conf, env)
	l = log.New("job")
	l.Infow("conf message", "conf", conf, "env", env)

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
		l.Panicf("队列服务异常退出: %+v", err)
	}
}
