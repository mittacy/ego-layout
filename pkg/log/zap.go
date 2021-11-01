package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var std = initStd()

func initStd() *zap.Logger {
	return NewWithLevel("default", globalLowLevel, zap.AddStacktrace(zapcore.WarnLevel))
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}

func Default() *zap.Logger {
	return std
}

// ResetDefault 重置默认日志文件
func ResetDefault(l *zap.Logger) {
	std = l
}

// New 创建新日志文件句柄
// @param logName 日志名
// @param opts 日志配置选项
// @return *Logger
func New(logName string, opts ...zap.Option) *zap.Logger {
	// 日志名检查
	logName = strings.TrimSpace(logName)
	if logName == "" {
		panic("the log file name is empty")
	}

	file, err := os.OpenFile(getLogPath(logName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return newWithWriter(file, zapcore.DebugLevel, opts...)
}

// NewWithLevel 创建新日志文件句柄,自定义最低日志级别
// @param logName 日志名
// @param level 最低日志级别
// @param opts 日志配置选项
// @return *Logger
func NewWithLevel(logName string, level zapcore.Level, opts ...zap.Option) *zap.Logger {
	// 日志名检查
	logName = strings.TrimSpace(logName)
	if logName == "" {
		panic("the log file name is empty")
	}

	file, err := os.OpenFile(getLogPath(logName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return newWithWriter(file, level, opts...)
}

func newWithWriter(writer io.Writer, level zapcore.Level, opts ...zap.Option) *zap.Logger {
	if writer == nil {
		panic("the writer is nil")
	}

	// 是否输出到控制台
	var syncer zapcore.WriteSyncer
	if logInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer))
	} else {
		syncer = zapcore.AddSync(writer)
	}

	// 配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "log_at",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "context",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02T15:04:05.000Z"))
		}, // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		syncer,
		level,
	)

	// 全局字段添加到每个日志中
	opts = append(opts, zap.Fields(globalFields...))

	return zap.New(core, opts...)
}

func getLogPath(name string) string {
	return fmt.Sprintf("%s/biz-%s.log", logPath, name)
}
