package loggerdto

import (
	"context"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type ErrDto struct {
	Ctx    context.Context
	Fields map[string]any
	Err    error
	Level  string
	File   string
	Func   string
	Line   int
}

func NewErr(ctx context.Context, level string, err error, fields map[string]any) *ErrDto {
	var fn string
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		fn = runtime.FuncForPC(pc).Name()
		fn = fn[strings.LastIndex(fn, "/")+1:]
	}

	return &ErrDto{
		Ctx:    ctx,
		Level:  level,
		Err:    err,
		Fields: fields,
		File:   file,
		Func:   fn,
		Line:   line,
	}
}

func (e *ErrDto) CallerFields() logrus.Fields {
	return logrus.Fields{
		"file": e.File,
		"func": e.Func,
		"line": e.Line,
	}
}
