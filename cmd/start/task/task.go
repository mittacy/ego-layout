package task

import (
	"github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/config"
	"github.com/mittacy/ego/library/log"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "task",
	Short: "start the scheduled task server",
	Long:  "start the scheduled task server. Example: server start task",
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
	bootstrap.InitTask(conf, env)
	l = log.New("task")

	var Tasks = config.Tasks(l)

	c := cron.New()

	for _, v := range Tasks {
		id, err := c.AddJob(v.Spec(), cron.NewChain(cron.Recover(&cronLog{l})).Then(v.Job()))
		if err != nil {
			l.Sugar().Errorf("task start fail, id: %d, jobName: %s, err: %s", id, v.Name(), err)
		} else {
			l.Sugar().Infof("task start success, id: %d, name: %s", id, v.Name())
		}
	}

	c.Start()
	defer c.Stop()
	select {}
}

type cronLog struct {
	l *log.Logger
}

// Info logs routine messages about cron's operation.
func (c *cronLog) Info(msg string, keysAndValues ...interface{}) {
	c.l.Infow(msg, keysAndValues...)
}

// Error logs an error condition.
func (c *cronLog) Error(err error, msg string, keysAndValues ...interface{}) {
	errPair := []interface{}{"err", err}
	keysAndValues = append(keysAndValues, errPair)
	c.l.Errorw(msg, keysAndValues...)
}
