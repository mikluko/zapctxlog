package ctxlog

import (
	"context"

	"go.uber.org/zap/zapcore"
)

// Debug level
func Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	Get(ctx).Debug(msg, fields...)
}

// Info level
func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	Get(ctx).Info(msg, fields...)
}

// Warn level
func Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	Get(ctx).Warn(msg, fields...)
}

// Error level
func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	Get(ctx).Error(msg, fields...)
}

// Fatal level
func Fatal(ctx context.Context, msg string, fields ...zapcore.Field) {
	Get(ctx).Fatal(msg, fields...)
}
