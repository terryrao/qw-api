package qwapi

import "context"

type Logger interface {
	Debug(ctx context.Context, str string, objs ...any)
	Info(ctx context.Context, str string, objjs ...any)
	Error(ctx context.Context, str string, objs ...any)
	Warn(ctx context.Context, str string, objs ...any)
}
