package log

import (
	"github.com/go-kratos/kratos/v2/log"
)

const (
	// LevelDebug is logger debug level.
	LevelDebug int8 = iota - 1
	// LevelInfo is logger info level.
	LevelInfo
	// LevelWarn is logger warn level.
	LevelWarn
	// LevelError is logger error level.
	LevelError
	// LevelFatal is logger fatal level
	LevelFatal
)

type Helper struct {
	log log.Logger
}

func NewHelper(logger log.Logger) *Helper {
	return &Helper{log: logger}
}

func (h *Helper) Log(level int8, info string, keyvals ...interface{}) {
	var kvs []interface{}
	kvs = append(kvs, "log_desc", info)
	kvs = append(kvs, keyvals...)
	_ = h.log.Log(log.Level(level), kvs...)
}

func (h *Helper) Debug(info string, keyvals ...interface{}) {
	h.Log(LevelDebug, info, keyvals...)
}

func (h *Helper) Info(info string, keyvals ...interface{}) {
	h.Log(LevelInfo, info, keyvals...)
}

func (h *Helper) Warn(info string, keyvals ...interface{}) {
	h.Log(LevelWarn, info, keyvals...)
}

func (h *Helper) Error(info string, keyvals ...interface{}) {
	h.Log(LevelError, info, keyvals...)
}

func (h *Helper) Fatal(info string, keyvals ...interface{}) {
	h.Log(LevelFatal, info, keyvals...)
}
