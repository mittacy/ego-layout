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
	return "0 8 * * ?"
}

func (t *Example) Job() cron.Job {
	return &Example{}
}

func (t *Example) Run() {
	// do something
	log.Info("Hello, this is the example task")
}

