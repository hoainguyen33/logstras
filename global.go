package logtras

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/hoainguyen33/logstras/log"
)

var global = &loggerAppliance{}

type loggerAppliance struct {
	lock sync.Mutex
	log.Logger
}

func init() {
	global.SetLogger(DefaultLogger)
}

func (a *loggerAppliance) SetLogger(in log.Logger) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.Logger = in
}

func (a *loggerAppliance) GetLogger() log.Logger {
	return a.Logger
}

func SetLogger(logger log.Logger) {
	global.SetLogger(logger)
}

func GetLogger() log.Logger {
	return global.GetLogger()
}

func Log(level log.Level, keyvals ...interface{}) {
	global.Log(level, keyvals...)
}

func Context(ctx context.Context) *Helper {
	return NewHelper(log.WithContext(ctx, global.Logger))
}

func Debug(a ...interface{}) {
	global.Log(log.LevelDebug, DefaultMessageKey, fmt.Sprint(a...))
}

func Debugf(format string, a ...interface{}) {
	global.Log(log.LevelDebug, DefaultMessageKey, fmt.Sprintf(format, a...))
}

func Debugw(keyvals ...interface{}) {
	global.Log(log.LevelDebug, keyvals...)
}

func Info(a ...interface{}) {
	global.Log(log.LevelInfo, DefaultMessageKey, fmt.Sprint(a...))
}

func Infof(format string, a ...interface{}) {
	global.Log(log.LevelInfo, DefaultMessageKey, fmt.Sprintf(format, a...))
}

func Infow(keyvals ...interface{}) {
	global.Log(log.LevelInfo, keyvals...)
}

func Warn(a ...interface{}) {
	global.Log(log.LevelWarn, DefaultMessageKey, fmt.Sprint(a...))
}

func Warnf(format string, a ...interface{}) {
	global.Log(log.LevelWarn, DefaultMessageKey, fmt.Sprintf(format, a...))
}

func Warnw(keyvals ...interface{}) {
	global.Log(log.LevelWarn, keyvals...)
}

func Error(a ...interface{}) {
	global.Log(log.LevelError, DefaultMessageKey, fmt.Sprint(a...))
}

func Errorf(format string, a ...interface{}) {
	global.Log(log.LevelError, DefaultMessageKey, fmt.Sprintf(format, a...))
}

func Errorw(keyvals ...interface{}) {
	global.Log(log.LevelError, keyvals...)
}

func Fatal(a ...interface{}) {
	global.Log(log.LevelFatal, DefaultMessageKey, fmt.Sprint(a...))
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	global.Log(log.LevelFatal, DefaultMessageKey, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func Fatalw(keyvals ...interface{}) {
	global.Log(log.LevelFatal, keyvals...)
	os.Exit(1)
}
