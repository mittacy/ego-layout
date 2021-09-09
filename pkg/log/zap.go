package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

type Level = zapcore.Level

const (
	DebugLevel  Level = zap.DebugLevel  // -1
	InfoLevel   Level = zap.InfoLevel   // 0
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3
	PanicLevel  Level = zap.PanicLevel  // 4
	FatalLevel  Level = zap.FatalLevel  // 5
)

type Field = zap.Field

func (l *Logger) GetZap() *zap.Logger {
	return l.l
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.l.DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.l.Sugar().Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	l.l.Sugar().Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.l.Sugar().Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.l.Sugar().Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.l.Sugar().DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.l.Sugar().Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.l.Sugar().Fatalf(template, args...)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *Logger) SugarDebug(args ...interface{}) {
	l.l.Sugar().Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *Logger) SugarInfo(args ...interface{}) {
	l.l.Sugar().Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *Logger) SugarWarn(args ...interface{}) {
	l.l.Sugar().Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *Logger) SugarError(args ...interface{}) {
	l.l.Sugar().Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) SugarDPanic(args ...interface{}) {
	l.l.Sugar().DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *Logger) SugarPanic(args ...interface{}) {
	l.l.Sugar().Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Logger) SugarFatal(args ...interface{}) {
	l.l.Sugar().Fatal(args...)
}

// function variables for all field types in github.com/uber-go/zap/field.go
var (
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Any         = zap.Any

	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Panic  = std.Panic
	Fatal  = std.Fatal
	Debug  = std.Debug

	Debugf  = std.Debugf
	Infof   = std.Infof
	Warnf   = std.Warnf
	Errorf  = std.Errorf
	DPanicf = std.DPanicf
	Panicf  = std.Panicf
	Fatalf  = std.Fatalf

	SugarDebug  = std.SugarDebug
	SugarInfo   = std.SugarInfo
	SugarWarn   = std.SugarWarn
	SugarError  = std.SugarError
	SugarDPanic = std.SugarDPanic
	SugarPanic  = std.SugarPanic
	SugarFatal  = std.SugarFatal
)

// ResetDefault 重置默认日志文件，不建议使用
func ResetDefault(l *Logger) {
	std = l
	Info = std.Info
	Warn = std.Warn
	Error = std.Error
	DPanic = std.DPanic
	Panic = std.Panic
	Fatal = std.Fatal
	Debug = std.Debug
	Debugf = std.Debugf
	Infof = std.Infof
	Warnf = std.Warnf
	Errorf = std.Errorf
	DPanicf = std.DPanicf
	Panicf = std.Panicf
	Fatalf = std.Fatalf
	SugarDebug = std.SugarDebug
	SugarInfo = std.SugarInfo
	SugarWarn = std.SugarWarn
	SugarError = std.SugarError
	SugarDPanic = std.SugarDPanic
	SugarPanic = std.SugarPanic
	SugarFatal = std.SugarFatal
}

type Logger struct {
	l     *zap.Logger // zap ensure that zap.Logger is safe for concurrent use
	level Level
}

var std *Logger

func Init(conf Config) {
	logConf = conf
	l := New("default")
	ResetDefault(l)
}

func Default() *Logger {
	return std
}

type Option = zap.Option

var (
	WithCaller    = zap.WithCaller
	AddCallerSkip = zap.AddCallerSkip
	AddStacktrace = zap.AddStacktrace
)

// New create a new logger
func New(logName string) *Logger {
	return NewWithLevel(logName, logConf.LowLevel)
}

func NewWithLevel(logName string, level Level) *Logger {
	if logName == "" {
		panic("the log file name is empty")
	}

	file, err := os.OpenFile(GetLogPath(logName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return newWithWriter(file, level, zap.Fields(zapcore.Field{
		Key:    "module_name",
		Type:   zapcore.StringType,
		String: logConf.ServerName,
	}), WithCaller(true), AddCallerSkip(1), AddStacktrace(zapcore.WarnLevel))
}

func newWithWriter(writer io.Writer, level Level, opts ...Option) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}

	// 是否输出到控制台
	var syncer zapcore.WriteSyncer
	if logConf.LogInConsole {
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
	logger := &Logger{
		l:     zap.New(core, opts...),
		level: level,
	}
	return logger
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}

func GetLogPath(name string) string {
	path := logConf.Path

	if path == "" {
		path = "./logs"
	}

	return fmt.Sprintf("%s/biz-%s.log", path, name)
}
