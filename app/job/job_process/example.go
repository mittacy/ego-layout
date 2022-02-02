package job_process

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego-layout/app/job/job_payload"
	"github.com/mittacy/ego/library/log"
)

// Processor 任务处理器
type ExampleProcessor struct {
	l *log.Logger
}

func NewExample() *ExampleProcessor {
	return &ExampleProcessor{
		l: log.New("example_job"),
	}
}

func (processor *ExampleProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p job_payload.ExamplePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// call service
	// service.Biz.Do()
	processor.l.Info("do something")

	return nil
}
