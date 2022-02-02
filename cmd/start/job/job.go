package job

import (
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/config"
	"github.com/mittacy/ego-layout/config/async_config"
	"github.com/mittacy/ego/library/async"
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

func run(cmd *cobra.Command, args []string) {
	bootstrap.InitJob(conf, env)
	async.InitLog()

	l = log.New("job")
	l.Infow("conf message", "conf", conf, "env", env)

	jobs := async_config.Jobs()
	srv := asynq.NewServer(config.AsyncRedisClientOpt(), config.AsyncConfig())

	mux := asynq.NewServeMux()
	for _, v := range jobs {
		mux.Handle(v.TypeName, v.Handler)
	}

	if err := srv.Run(mux); err != nil {
		l.Panicf("队列服务异常退出: %+v", err)
	}
}
