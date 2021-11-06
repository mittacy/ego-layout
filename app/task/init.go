package task

import (
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/robfig/cron/v3"
)

type Task interface {
	// Name 任务名
	Name() string
	// Spec 定时规则
	Spec() string
	// Job cron任务
	Job() cron.Job
}

func Tasks(logger *log.Logger) []Task {
	return []Task{
		NewExampleTask(logger),
	}
}

func StartTasks() {
	l := log.New("task")

	var Tasks = Tasks(l)

	c := cron.New()

	for _, v := range Tasks {
		id, err := c.AddJob(v.Spec(), v.Job())
		if err != nil {
			l.Sugar().Errorf("task start fail, id: %d, jobName: %s, err: %s", id, v.Name(), err)
		} else {
			l.Sugar().Infof("task start success, id: %d, name: %s", id, v.Name())
		}
	}

	c.Start()
}
