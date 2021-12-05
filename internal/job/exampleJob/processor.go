package exampleJob

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego-layout/pkg/async"
	"github.com/mittacy/ego-layout/pkg/log"
)

func NewProcessor() *Processor {
	return &Processor{
		l: async.GetLogger(),
	}
}

// Processor 任务处理器, 实现 asynq.Handler 接口
type Processor struct {
	l *log.Logger
}

func (processor *Processor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p Payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// do work...
	processor.l.Infof("数据: %+v", p)

	return nil
}
