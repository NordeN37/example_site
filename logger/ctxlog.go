package logger

import (
	"context"
	"github.com/rs/zerolog"
)

const (
	loggerCtxKey = "loggerCtx"
)

// GetFromCtx достаёт из контекста логгер
func GetFromCtx(ctx context.Context) *zerolog.Logger {
	return ctx.Value(loggerCtxKey).(*zerolog.Logger)
}

// AddToCtx добавляет в контекст логгер
func AddToCtx(ctx *context.Context, logger *zerolog.Logger) {
	*ctx = context.WithValue(*ctx, loggerCtxKey, logger)
}
