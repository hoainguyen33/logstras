package log

import (
	"bytes"
	"io"
	"log"
	"sync"
)

type stdLogger struct {
	log    *log.Logger
	pool   *sync.Pool
	format Format
}

func NewStdLogger(w io.Writer) Logger {
	return &stdLogger{
		log: log.New(w, "", 0),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

func (l *stdLogger) Log(level Level, keyvals ...interface{}) {
	if len(keyvals) == 0 {
		return
	}
	buf := l.pool.Get().(*bytes.Buffer)
	l.format.Format(buf, level, keyvals...)
	l.log.Output(4, buf.String())
	buf.Reset()
	l.pool.Put(buf)
}

func (l *stdLogger) Close() error {
	return nil
}
