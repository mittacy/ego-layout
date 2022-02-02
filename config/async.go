package config

import (
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego/library/async"
	"github.com/mittacy/ego/library/log"
	"github.com/spf13/viper"
	"time"
)

// AsyncRedisClientOpt is used to create a redis client that connects
// to a redis server directly.
func AsyncRedisClientOpt() asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Network:      viper.GetString("ASYNC_Network"),
		Addr:         viper.GetString("ASYNC_Addr"),
		Username:     viper.GetString("ASYNC_Username"),
		Password:     viper.GetString("ASYNC_Password"),
		DB:           viper.GetInt("ASYNC_DB"),
		DialTimeout:  viper.GetDuration("ASYNC_DialTimeout") * time.Second,
		ReadTimeout:  viper.GetDuration("ASYNC_ReadTimeout") * time.Second,
		WriteTimeout: viper.GetDuration("ASYNC_WriteTimeout") * time.Second,
		PoolSize:     viper.GetInt("ASYNC_PoolSize"),
	}
}

// AsyncConfig specifies the server's background-task processing behavior.
func AsyncConfig() asynq.Config {
	return asynq.Config{
		Concurrency: viper.GetInt("ASYNC_Concurrency"),
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		ErrorHandler:        nil,
		Logger:              async.NewLogger(log.New("job")),
		LogLevel:            asynq.InfoLevel,
		ShutdownTimeout:     time.Second * 10, // 关闭服务时等待等待进行中的任务
		HealthCheckFunc:     nil,
		HealthCheckInterval: 0,
	}
}
