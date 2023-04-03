package logstras

import "github.com/hoainguyen33/logstras/log"

type LoggerHook interface {
	Log(args ...interface{})
}

type loggerHook struct {
	log LoggerHook
}

func (l *loggerHook) Log(level log.Level, keyvals ...interface{}) {
	l.log.Log(level, keyvals)
}

func NewLoggerHook(log LoggerHook) log.Logger {
	return &loggerHook{log: log}
}
