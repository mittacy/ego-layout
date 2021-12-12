package exampleJobTask

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego-layout/pkg/async"
)

// TypeName 任务名(业务名:操作名)
const TypeName = "bizName:actionName"

// Payload 执行任务需要的数据
type Payload struct {
	// 异步处理时需要的数据
}

// NewTask 新建任务
func NewTask(opts ...asynq.Option) (*asynq.Task, error) {
	data := Payload{}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeName, payload, opts...), nil
}

// NewTaskAndEnqueue 新建任务并加入任务队列
func NewTaskAndEnqueue(opts ...asynq.Option) error {
	task, err := NewTask()
	if err != nil {
		return err
	}

	return async.Enqueue(task, opts...)
}
