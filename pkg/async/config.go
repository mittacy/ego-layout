package async

import (
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego/library/log"
	"github.com/spf13/viper"
	"sync"
	"time"
)

var (
	logger   *log.Logger
	client   *asynq.Client
	initOnce sync.Once
)

func initDependency() {
	logger = log.New("job")
	client = asynq.NewClient(GetDefaultRedisConnOpt())
}

// GetDefaultClient 获取默认连接
func GetDefaultClient() *asynq.Client {
	initOnce.Do(initDependency)

	return client
}

// GetLogger 获取异步任务日志文件
func GetLogger() *log.Logger {
	initOnce.Do(initDependency)

	return logger
}

// GetDefaultRedisConnOpt 获取默认redis配置
func GetDefaultRedisConnOpt() asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Network:      viper.GetString("QSYNC_Network"),
		Addr:         viper.GetString("QSYNC_Addr"),
		Username:     viper.GetString("QSYNC_Username"),
		Password:     viper.GetString("QSYNC_Password"),
		DB:           viper.GetInt("QSYNC_DB"),
		DialTimeout:  viper.GetDuration("QSYNC_DialTimeout") * time.Second,
		ReadTimeout:  viper.GetDuration("QSYNC_ReadTimeout") * time.Second,
		WriteTimeout: viper.GetDuration("QSYNC_WriteTimeout") * time.Second,
		PoolSize:     viper.GetInt("QSYNC_PoolSize"),
	}
}

// GetDefaultServerConfig 获取默认异步任务服务配置
func GetDefaultServerConfig() asynq.Config {
	initOnce.Do(initDependency)

	return asynq.Config{
		Concurrency: viper.GetInt("QSYNC_Concurrency"),
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		ErrorHandler:        nil,
		Logger:              NewLogger(logger),
		LogLevel:            asynq.InfoLevel,
		ShutdownTimeout:     time.Second * 10, // 关闭服务时等待等待进行中的任务
		HealthCheckFunc:     nil,
		HealthCheckInterval: 0,
	}
}
