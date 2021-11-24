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

type Logger struct {
	l *zap.Logger
}

func (l *Logger) GetZap() *zap.Logger {
	return l.l
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

var (
	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Panic  = std.Panic
	Fatal  = std.Fatal
	Debug  = std.Debug
	Sugar  = std.Sugar

	Infof   = std.Infof
	Warnf   = std.Warnf
	Errorf  = std.Errorf
	DPanicf = std.DPanicf
	Panicf  = std.Panicf
	Fatalf  = std.Fatalf
	Debugf  = std.Debugf

	Infow   = std.Infow
	Warnw   = std.Warnw
	Errorw  = std.Errorw
	DPanicw = std.DPanicw
	Panicw  = std.Panicw
	Fatalw  = std.Fatalw
	Debugw  = std.Debugw

	CopierErrLog      = std.CopierErrLog
	TransformErrLog   = std.TransformErrLog
	JsonMarshalErrLog = std.JsonMarshalErrLog
	CacheErrLog       = std.CacheErrLog
	MysqlErrLog       = std.MysqlErrLog
	BizErrorLog            = std.BizErrorLog
)

var std = initStd()

func initStd() *Logger {
	return NewWithLevel("default", globalLowLevel, zap.AddStacktrace(zapcore.WarnLevel))
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}

func Default() *Logger {
	return std
}

// ResetDefault 重置默认日志文件
func ResetDefault(l *Logger) {
	std = l
	Sugar = std.Sugar
	Info = std.Info
	Warn = std.Warn
	Error = std.Error
	DPanic = std.DPanic
	Panic = std.Panic
	Fatal = std.Fatal
	Debug = std.Debug

	Infof = std.Infof
	Warnf = std.Warnf
	Errorf = std.Errorf
	DPanicf = std.DPanicf
	Panicf = std.Panicf
	Fatalf = std.Fatalf
	Debugf = std.Debugf

	Infow = std.Infow
	Warnw = std.Warnw
	Errorw = std.Errorw
	DPanicw = std.DPanicw
	Panicw = std.Panicw
	Fatalw = std.Fatalw
	Debugw = std.Debugw

	CopierErrLog = std.CopierErrLog
	TransformErrLog = std.TransformErrLog
	JsonMarshalErrLog = std.JsonMarshalErrLog
	CacheErrLog = std.CacheErrLog
	MysqlErrLog = std.MysqlErrLog
	BizErrorLog = std.BizErrorLog
}

// New 创建新日志文件句柄
// @param logName 日志名
// @param opts 日志配置选项
// @return *Logger
func New(logName string, opts ...zap.Option) *Logger {
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
func NewWithLevel(logName string, level zapcore.Level, opts ...zap.Option) *Logger {
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

func newWithWriter(writer io.Writer, level zapcore.Level, opts ...zap.Option) *Logger {
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
		TimeKey:  "log_at",
		LevelKey: "level",
		NameKey:  "logger",
		//CallerKey:     "caller",
		MessageKey: "context",
		//StacktraceKey: "stacktrace",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}, // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		//EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName: zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if !jsonEncoder {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		syncer,
		level,
	)

	// 全局字段添加到每个日志中
	opts = append(opts, zap.Fields(globalFields...))

	logger := &Logger{
		l: zap.New(core, opts...),
	}

	return logger
}

func getLogPath(name string) string {
	return fmt.Sprintf("%s/biz-%s.log", logPath, name)
}
