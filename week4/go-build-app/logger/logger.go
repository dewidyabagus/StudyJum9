package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap *zap.Logger
}

var (
	String = zap.String
)

type Field = zapcore.Field

func New(lvl string, output ...string) (*Logger, error) {
	level := zap.InfoLevel
	switch lvl {
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	}

	if len(output) == 0 {
		output = []string{"stdout"}
	}

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: output,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "lvl",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	zapLog, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("build logger failed: %w", err)
	}

	return &Logger{zap: zapLog}, nil
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.zap.Info(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.zap.Error(msg, fields...)
}
