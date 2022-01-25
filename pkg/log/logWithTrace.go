package log

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

const (
	RequestId = "requestId"
)

func GetRequestId(c *gin.Context) string {
	return c.GetString(RequestId)
}

// DebugWithTrace logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) DebugWithTrace(c *gin.Context, msg string, fields ...zapcore.Field) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	fields = append(fields, field)
	l.l.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) InfoWithTrace(c *gin.Context, msg string, fields ...zapcore.Field) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	fields = append(fields, field)
	l.l.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) WarnWithTrace(c *gin.Context, msg string, fields ...zapcore.Field) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	fields = append(fields, field)
	l.l.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) ErrorWithTrace(c *gin.Context, msg string, fields ...zapcore.Field) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	fields = append(fields, field)
	l.l.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func (l *Logger) DPanicWithTrace(c *gin.Context, msg string, fields ...zapcore.Field) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	fields = append(fields, field)
	l.l.DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func (l *Logger) PanicWithTrace(c *gin.Context, msg string, fields ...zapcore.Field) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	fields = append(fields, field)
	l.l.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (l *Logger) FatalWithTrace(c *gin.Context, msg string, fields ...zapcore.Field) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	fields = append(fields, field)
	l.l.Fatal(msg, fields...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) DebugfWithTrace(c *gin.Context, template string, args ...interface{}) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	l.Sugar().With(field).Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) InfofWithTrace(c *gin.Context, template string, args ...interface{}) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	l.Sugar().With(field).Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) WarnfWithTrace(c *gin.Context, template string, args ...interface{}) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	l.Sugar().With(field).Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) ErrorfWithTrace(c *gin.Context, template string, args ...interface{}) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	l.Sugar().With(field).Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) DPanicfWithTrace(c *gin.Context, template string, args ...interface{}) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	l.Sugar().With(field).DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *Logger) PanicfWithTrace(c *gin.Context, template string, args ...interface{}) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	l.Sugar().With(field).Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Logger) FatalfWithTrace(c *gin.Context, template string, args ...interface{}) {
	field := zapcore.Field{
		Key:    RequestId,
		Type:   zapcore.StringType,
		String: GetRequestId(c),
	}
	l.Sugar().With(field).Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func (l *Logger) DebugwWithTrace(c *gin.Context, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, RequestId)
	keysAndValues = append(keysAndValues, GetRequestId(c))
	l.Sugar().Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) InfowWithTrace(c *gin.Context, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, RequestId)
	keysAndValues = append(keysAndValues, GetRequestId(c))
	l.Sugar().Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) WarnwWithTrace(c *gin.Context, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, RequestId)
	keysAndValues = append(keysAndValues, GetRequestId(c))
	l.Sugar().Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) ErrorwWithTrace(c *gin.Context, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, RequestId)
	keysAndValues = append(keysAndValues, GetRequestId(c))
	l.Sugar().Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) DPanicwWithTrace(c *gin.Context, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, RequestId)
	keysAndValues = append(keysAndValues, GetRequestId(c))
	l.Sugar().DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (l *Logger) PanicwWithTrace(c *gin.Context, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, RequestId)
	keysAndValues = append(keysAndValues, GetRequestId(c))
	l.Sugar().Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (l *Logger) FatalwWithTrace(c *gin.Context, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, RequestId)
	keysAndValues = append(keysAndValues, GetRequestId(c))
	l.Sugar().Fatalw(msg, keysAndValues...)
}
