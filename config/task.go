package config

import (
	"github.com/mittacy/ego-layout/app/task"
	"github.com/mittacy/ego/library/log"
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
		task.NewExample(logger),
	}
}
