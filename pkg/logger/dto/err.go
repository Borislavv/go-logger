package loggerdto

import (
	"context"
	"github.com/Borislavv/go-logger/pkg/logger"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type ErrDto struct {
	Ctx    context.Context
	Fields logger.Fields
	Err    error
	Level  string
	File   string
	Func   string
	Line   int
}

func NewErr(ctx context.Context, level string, err error, fields logger.Fields) *ErrDto {
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
