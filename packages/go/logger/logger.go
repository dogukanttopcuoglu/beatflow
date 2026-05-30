package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	Environment string
	ServiceName string
	Level       string
}

func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return zapcore.DebugLevel

	case "warn", "warning":
		return zapcore.WarnLevel

	case "error":
		return zapcore.ErrorLevel

	default:
		return zapcore.InfoLevel
	}
}

func New(opts Options) (*zap.Logger, error) {
	level := parseLevel(opts.Level)

	var cfg zap.Config

	switch strings.ToLower(strings.TrimSpace(opts.Environment)) {
	case "development", "dev", "local":
		cfg = zap.NewDevelopmentConfig()
		cfg.Level = zap.NewAtomicLevelAt(level)
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05")
	default:
		cfg = productionConfig(level)
	}

	log, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	if opts.ServiceName != "" {
		log = log.With(zap.String("service_name", opts.ServiceName))
	}

	return log, nil
}

func productionConfig(level zapcore.Level) zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
}
