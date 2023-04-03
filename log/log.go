package log

import (
	"context"
)

type Logger interface {
	Log(level Level, keyvals ...interface{})
}

type logger struct {
	formater Formater
	logger   Logger
	ctx      context.Context
}

func (c *logger) Log(level Level, keyvals ...interface{}) {
	var kvs []interface{}
	if len(keyvals) == 1 {
		kvs = c.formater.Bind(keyvals[0])
	} else {
		kvs = c.formater.Bind(keyvals)
	}
	if c.formater.HasValuer(kvs) {
		bindValues(c.ctx, kvs)
	}
	c.logger.Log(level, kvs...)
}

func With(l Logger, kv ...interface{}) Logger {
	var formater Formater
	if len(kv) == 1 {
		formater = NewFormater(kv[0])
	} else {
		formater = NewFormater(kv)
	}
	c, ok := l.(*logger)
	if !ok {
		return &logger{logger: l, ctx: context.Background(), formater: formater}
	}
	return &logger{
		logger:   c.logger,
		ctx:      c.ctx,
		formater: formater,
	}
}

func WithContext(ctx context.Context, l Logger) Logger {
	c, ok := l.(*logger)
	if !ok {
		return &logger{logger: l, ctx: ctx, formater: NewFormater(nil)}
	}
	return &logger{
		logger:   c.logger,
		ctx:      ctx,
		formater: c.formater,
	}
}
