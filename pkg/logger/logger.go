package logger

import (
	"context"
	"io"
)

type Fields = map[string]any

type Logger interface {
	DebugMsg(ctx context.Context, msg string, fields Fields)
	InfoMsg(ctx context.Context, msg string, fields Fields)
	WarningMsg(ctx context.Context, msg string, fields Fields)
	ErrorMsg(ctx context.Context, msg string, fields Fields)
	FatalMsg(ctx context.Context, msg string, fields Fields)
	PanicMsg(ctx context.Context, msg string, fields Fields)

	Debug(ctx context.Context, err error, fields Fields) error
	Info(ctx context.Context, err error, fields Fields) error
	Warning(ctx context.Context, err error, fields Fields) error
	Error(ctx context.Context, err error, fields Fields) error
	Fatal(ctx context.Context, err error, fields Fields) error
	Panic(ctx context.Context, err error, fields Fields) error
}

type CancelFunc = func()

type Outputer interface {
	io.Closer
	io.Writer
}
