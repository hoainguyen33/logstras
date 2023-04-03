package log

import (
	"context"
	"fmt"
	"os"
)

var DefaultMessageKey = "payload"

type Option func(*Helper)

type Helper struct {
	logger Logger
	msgKey string
}

func WithMessageKey(k string) Option {
	return func(opts *Helper) {
		opts.msgKey = k
	}
}

func NewHelper(logger Logger, opts ...Option) *Helper {
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
		logger: WithContext(ctx, h.logger),
	}
}

func (h *Helper) Log(level Level, keyvals ...interface{}) {
	h.logger.Log(level, keyvals...)
}

func (h *Helper) Debug(a ...interface{}) {
	as := append([]interface{}{h.msgKey}, a...)
	h.logger.Log(LevelDebug, h.msgKey, as)
}

func (h *Helper) Debugf(format string, a ...interface{}) {
	h.logger.Log(LevelDebug, h.msgKey, fmt.Sprintf(format, a...))
}

func (h *Helper) Debugw(keyvals ...interface{}) {
	h.logger.Log(LevelDebug, keyvals...)
}

func (h *Helper) Info(a ...interface{}) {
	as := append([]interface{}{h.msgKey}, a...)
	h.logger.Log(LevelInfo, as...)
}

func (h *Helper) Infof(format string, a ...interface{}) {
	h.logger.Log(LevelInfo, h.msgKey, fmt.Sprintf(format, a...))
}

func (h *Helper) Infow(keyvals ...interface{}) {
	h.logger.Log(LevelInfo, keyvals...)
}

func (h *Helper) Warn(a ...interface{}) {
	as := append([]interface{}{h.msgKey}, a...)
	h.logger.Log(LevelWarn, as)
}

func (h *Helper) Warnf(format string, a ...interface{}) {
	h.logger.Log(LevelWarn, h.msgKey, fmt.Sprintf(format, a...))
}

func (h *Helper) Warnw(keyvals ...interface{}) {
	h.logger.Log(LevelWarn, keyvals...)
}

func (h *Helper) Error(a ...interface{}) {
	fmt.Println(222, a)
	as := append([]interface{}{h.msgKey}, a...)
	h.logger.Log(LevelError, as...)
}

func (h *Helper) Errorf(format string, a ...interface{}) {
	fmt.Println(222, a)
	h.logger.Log(LevelError, h.msgKey, fmt.Sprintf(format, a...))
}

func (h *Helper) Errorw(keyvals ...interface{}) {
	fmt.Println(222, keyvals)
	h.logger.Log(LevelError, keyvals...)
}

func (h *Helper) Fatal(a ...interface{}) {
	as := append([]interface{}{h.msgKey}, a...)
	h.logger.Log(LevelFatal, as)
	os.Exit(1)
}

func (h *Helper) Fatalf(format string, a ...interface{}) {
	h.logger.Log(LevelFatal, h.msgKey, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func (h *Helper) Fatalw(keyvals ...interface{}) {
	h.logger.Log(LevelFatal, keyvals...)
	os.Exit(1)
}
