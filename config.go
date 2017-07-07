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

type Option func(*zap.Config)

func NewConfig(level string, opts ...Option) *zap.Config {
	cfg := zap.Config{
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
		OutputPaths:       []string{},
		ErrorOutputPaths:  []string{},
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	// Set default output and error paths if not set by options
	if len(cfg.OutputPaths) == 0 {
		cfg.OutputPaths = []string{"stdout"}
	}
	if len(cfg.ErrorOutputPaths) == 0 {
		cfg.OutputPaths = []string{"stderr"}
	}

	return &cfg
}

func OutputPath(s string) Option {
	return func(cfg *zap.Config) {
		cfg.OutputPaths = append(cfg.OutputPaths, s)
	}
}

func ErrorOutputPath(s string) Option {
	return func(cfg *zap.Config) {
		cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, s)
	}
}

func Development() Option {
	return func(cfg *zap.Config) {
		cfg.Development = true
	}
}

func EnableCaller() Option {
	return func(cfg *zap.Config) {
		cfg.DisableCaller = false
	}
}

func EnableStacktrace() Option {
	return func(cfg *zap.Config) {
		cfg.DisableStacktrace = false
	}
}
