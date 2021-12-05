package async

import (
	"fmt"
	"github.com/hibiken/asynq"
)

// Enqueue 加入任务队列
// @param task 任务
// @param opts 可选配置，失败重试、指定队列、优先级……
// @return error
func Enqueue(task *asynq.Task, opts ...asynq.Option) error {
	initOnce.Do(initDependency)

	// 加入队列
	info, err := client.Enqueue(task, opts...)
	if err != nil {
		msg := fmt.Sprintf("send task fail, type: %s", task.Type())
		logger.Errorw(msg, "payload", task.Payload(), "err", err)
		return err
	}

	logger.Infof("%s, taskId: %s 加入 %s 队列成功", task.Type(), info.ID, info.Queue)
	return nil
}
