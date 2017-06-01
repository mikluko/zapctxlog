package ctxlog

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// String field
func String(k, v string) zapcore.Field {
	return zap.String(k, v)
}

// Strings field
func Strings(k string, v []string) zapcore.Field {
	return zap.Strings(k, v)
}

// Stringer field
func Stringer(k string, v fmt.Stringer) zapcore.Field {
	return zap.Stringer(k, v)
}

// Err field
func Err(e error) zapcore.Field {
	return zap.Error(e)
}

// Duration field
func Duration(k string, v time.Duration) zapcore.Field {
	return zap.Duration(k, v)
}

// Bool field
func Bool(k string, v bool) zapcore.Field {
	return zap.Bool(k, v)
}

// Int field
func Int(k string, v int) zapcore.Field {
	return zap.Int(k, v)
}
