package log

import (
	"context"
	"fmt"
	"os"
	"sync"
)

var global = &loggerAppliance{}

type loggerAppliance struct {
	lock sync.Mutex
	Logger
}

func init() {
	global.SetLogger(DefaultLogger)
}

func (a *loggerAppliance) SetLogger(in Logger) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.Logger = in
}

func (a *loggerAppliance) GetLogger() Logger {
	return a.Logger
}

func SetLogger(logger Logger) {
	global.SetLogger(logger)
}

func GetLogger() Logger {
	return global.GetLogger()
}

func Log(level Level, keyvals ...interface{}) {
	global.Log(level, keyvals...)
}

func Context(ctx context.Context) *Helper {
	return NewHelper(WithContext(ctx, global.Logger))
}

func Debug(a ...interface{}) {
	global.Log(LevelDebug, DefaultMessageKey, fmt.Sprint(a...))
}

func Debugf(format string, a ...interface{}) {
	global.Log(LevelDebug, DefaultMessageKey, fmt.Sprintf(format, a...))
}

func Debugw(keyvals ...interface{}) {
	global.Log(LevelDebug, keyvals...)
}

func Info(a ...interface{}) {
	global.Log(LevelInfo, DefaultMessageKey, fmt.Sprint(a...))
}

func Infof(format string, a ...interface{}) {
	global.Log(LevelInfo, DefaultMessageKey, fmt.Sprintf(format, a...))
}

func Infow(keyvals ...interface{}) {
	global.Log(LevelInfo, keyvals...)
}

func Warn(a ...interface{}) {
	global.Log(LevelWarn, DefaultMessageKey, fmt.Sprint(a...))
}

func Warnf(format string, a ...interface{}) {
	global.Log(LevelWarn, DefaultMessageKey, fmt.Sprintf(format, a...))
}

func Warnw(keyvals ...interface{}) {
	global.Log(LevelWarn, keyvals...)
}

func Error(a ...interface{}) {
	global.Log(LevelError, DefaultMessageKey, fmt.Sprint(a...))
}

func Errorf(format string, a ...interface{}) {
	global.Log(LevelError, DefaultMessageKey, fmt.Sprintf(format, a...))
}

func Errorw(keyvals ...interface{}) {
	global.Log(LevelError, keyvals...)
}

func Fatal(a ...interface{}) {
	global.Log(LevelFatal, DefaultMessageKey, fmt.Sprint(a...))
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	global.Log(LevelFatal, DefaultMessageKey, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func Fatalw(keyvals ...interface{}) {
	global.Log(LevelFatal, keyvals...)
	os.Exit(1)
}
