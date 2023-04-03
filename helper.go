package logstras

import (
	"context"
	"fmt"

	"os"

	"github.com/hoainguyen33/logstras/log"
)

var DefaultMessageKey = "payload"

type Option func(*Helper)

type Helper struct {
	logger log.Logger
	msgKey string
}

func WithMessageKey(k string) Option {
	return func(opts *Helper) {
		opts.msgKey = k
	}
}

func NewHelper(logger log.Logger, opts ...Option) *Helper {
	options := &Helper{
		msgKey: DefaultMessageKey,
		logger: logger,
	}
	for _, o := range opts {
		o(options)
	}
	return options
}

func (h *Helper) WithContext(ctx context.Context) *Helper {
	return &Helper{
		msgKey: h.msgKey,
		logger: log.WithContext(ctx, h.logger),
	}
}

func (h *Helper) Log(level log.Level, keyvals ...interface{}) {
	h.logger.Log(level, keyvals...)
}

func (h *Helper) Debug(a ...interface{}) {
	as := append([]interface{}{h.msgKey}, a...)
	h.logger.Log(log.LevelDebug, h.msgKey, as)
}

func (h *Helper) Debugf(format string, a ...interface{}) {
	h.logger.Log(log.LevelDebug, h.msgKey, fmt.Sprintf(format, a...))
}

func (h *Helper) Debugw(keyvals ...interface{}) {
	h.logger.Log(log.LevelDebug, keyvals...)
}

func (h *Helper) Info(a ...interface{}) {
	as := append([]interface{}{h.msgKey}, a...)
	h.logger.Log(log.LevelInfo, as...)
}

func (h *Helper) Infof(format string, a ...interface{}) {
	h.logger.Log(log.LevelInfo, h.msgKey, fmt.Sprintf(format, a...))
}

func (h *Helper) Infow(keyvals ...interface{}) {
	h.logger.Log(log.LevelInfo, keyvals...)
}

func (h *Helper) Warn(a ...interface{}) {
	as := append([]interface{}{h.msgKey}, a...)
	h.logger.Log(log.LevelWarn, as)
}

func (h *Helper) Warnf(format string, a ...interface{}) {
	h.logger.Log(log.LevelWarn, h.msgKey, fmt.Sprintf(format, a...))
}

func (h *Helper) Warnw(keyvals ...interface{}) {
	h.logger.Log(log.LevelWarn, keyvals...)
}

func (h *Helper) Error(a ...interface{}) {
	as := append([]interface{}{h.msgKey}, a...)
	h.logger.Log(log.LevelError, as...)
}

func (h *Helper) Errorf(format string, a ...interface{}) {
	h.logger.Log(log.LevelError, h.msgKey, fmt.Sprintf(format, a...))
}

func (h *Helper) Errorw(keyvals ...interface{}) {
	h.logger.Log(log.LevelError, keyvals...)
}

func (h *Helper) Fatal(a ...interface{}) {
	as := append([]interface{}{h.msgKey}, a...)
	h.logger.Log(log.LevelFatal, as)
	os.Exit(1)
}

func (h *Helper) Fatalf(format string, a ...interface{}) {
	h.logger.Log(log.LevelFatal, h.msgKey, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func (h *Helper) Fatalw(keyvals ...interface{}) {
	h.logger.Log(log.LevelFatal, keyvals...)
	os.Exit(1)
}
