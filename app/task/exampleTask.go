package task

import (
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/robfig/cron/v3"
)

type ExampleTask struct {
	logger *log.Logger
}

func NewExampleTask(logger *log.Logger) *ExampleTask {
	return &ExampleTask{logger: logger}
}

func (t *ExampleTask) Name() string {
	return "exampleTask"
}

func (t *ExampleTask) Spec() string {
	return "0 8 * * ?"
}

func (t *ExampleTask) Job() cron.Job {
	return &ExampleTask{}
}

func (t *ExampleTask) Run() {
	// do something
	log.Info("Hello, this is the example task")
}

