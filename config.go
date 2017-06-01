package ctxlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LevelFromString(s string) (zap.AtomicLevel, error) {
	l := zap.NewAtomicLevel()
	err := l.UnmarshalText([]byte(s))
	return l, err
}

func MustLevelFromString(s string) zap.AtomicLevel {
	l, err := LevelFromString(s)
	if err != nil {
		panic(err)
	}
	return l
}

func NewConfig(level string) *zap.Config {
	return &zap.Config{
		Level: MustLevelFromString(level),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:       "level",
			MessageKey:     "msg",
			NameKey:        "name",
			StacktraceKey:  "trace",
			CallerKey:      "caller",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
	}
}
