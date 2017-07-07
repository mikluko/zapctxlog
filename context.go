package ctxlog

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey uint

const (
	ctxLogger ctxKey = iota + 1
)

// Get gets a logger from the context
func Get(ctx context.Context) *zap.Logger {
	var (
		logger *zap.Logger
		ok     bool
	)
	logger, ok = ctx.Value(ctxLogger).(*zap.Logger)
	if !ok {
		logger = zap.NewNop()
	}
	return logger
}

// Push pushes a logger to the context
func Push(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger, logger)
}

// Named appends name to the logger in the context
func Named(ctx context.Context, name string) context.Context {
	return Push(ctx, Get(ctx).Named(name))
}

// With pushes log fields to the context
func With(ctx context.Context, fields ...zapcore.Field) context.Context {
	return Push(ctx, Get(ctx).With(fields...))
}

// Logger initializes a logger and pushes it into the context
func Logger(ctx context.Context, name string, opts ...Option) context.Context {
	logger, err := NewConfig(opts...).Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return Push(ctx, logger.Named(name))
}
