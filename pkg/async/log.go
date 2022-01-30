package async

import (
	"github.com/hibiken/asynq"
	"github.com/mittacy/ego/library/log"
)

type asynqLogger struct {
	l *log.Logger
}

func NewLogger(l *log.Logger) asynq.Logger {
	return &asynqLogger{l: l}
}

// Debug logs a message at Debug level.
func (l *asynqLogger) Debug(args ...interface{}) {
	l.l.Sugar().Debug(args...)
}

// Info logs a message at Info level.
func (l *asynqLogger) Info(args ...interface{}) {
	l.l.Sugar().Info(args...)
}

// Warn logs a message at Warning level.
func (l *asynqLogger) Warn(args ...interface{}) {
	l.l.Sugar().Warn(args...)
}

// Error logs a message at Error level.
func (l *asynqLogger) Error(args ...interface{}) {
	l.l.Sugar().Error(args...)
}

// Fatal logs a message at Fatal level
// and process will exit with status set to 1.
func (l *asynqLogger) Fatal(args ...interface{}) {
	l.l.Sugar().Fatal(args...)
}
