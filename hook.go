package log

type LoggerHook interface {
	Log(args ...interface{})
}

type loggerHook struct {
	log LoggerHook
}

func (l *loggerHook) Log(level Level, keyvals ...interface{}) {
	l.log.Log(level, keyvals)
}

func NewLoggerHook(log LoggerHook) Logger {
	return &loggerHook{log: log}
}
