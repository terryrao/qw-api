package qwapi

import "context"

type Logger interface {
	Debug(ctx context.Context, str string, objs ...any)
	Info(ctx context.Context, str string, objjs ...any)
	Error(ctx context.Context, str string, objs ...any)
	Warn(ctx context.Context, str string, objs ...any)
}

type emptyLogger struct {
}

// mp *emptyLogger

func (e *emptyLogger) Debug(ctx context.Context, str string, objs ...any) {}
func (e *emptyLogger) Info(ctx context.Context, str string, objjs ...any) {}
func (e *emptyLogger) Error(ctx context.Context, str string, objs ...any) {}
func (e *emptyLogger) Warn(ctx context.Context, str string, objs ...any)  {}
