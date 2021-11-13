package task

import (
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/robfig/cron/v3"
)

type Example struct {
	logger *log.Logger
}

func NewExample(logger *log.Logger) *Example {
	return &Example{logger: logger}
}

func (t *Example) Name() string {
	return "Example"
}

func (t *Example) Spec() string {
	return "@every 10s"
}

func (t *Example) Job() cron.Job {
	return t
}

func (t *Example) Run() {
	// do something
	t.logger.Infow("Hello, this is the example task", "task", t.Name())
}

