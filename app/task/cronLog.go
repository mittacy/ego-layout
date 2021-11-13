package task

import "github.com/mittacy/ego-layout/pkg/log"

type cronLog struct {
	l *log.Logger
}

// Info logs routine messages about cron's operation.
func (c *cronLog) Info(msg string, keysAndValues ...interface{}) {
	c.l.Infow(msg, keysAndValues...)
}

// Error logs an error condition.
func (c *cronLog) Error(err error, msg string, keysAndValues ...interface{}) {
	errPair := []interface{}{"err", err}
	keysAndValues = append(errPair, keysAndValues...)
	c.l.Errorw(msg, keysAndValues...)
}

