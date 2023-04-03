package logtras

import (
	"bytes"
	"io"
	l "log"
	"sync"

	"github.com/hoainguyen33/logstras/log"
)

var DefaultLogger log.Logger = NewStdLogger(l.Writer())

type stdLogger struct {
	log    *l.Logger
	pool   *sync.Pool
	format log.Format
}

func NewStdLogger(w io.Writer) log.Logger {
	return &stdLogger{
		log: l.New(w, "", 0),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

func (l *stdLogger) Log(level log.Level, keyvals ...interface{}) {
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
