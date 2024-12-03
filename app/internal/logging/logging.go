package logging

import (
	"context"
	"log/slog"
	"os"
	"runtime"

	"event_system/internal/config"
)

type Logger interface {
	Debug(msg string, attrs ...any)
	Info(msg string, attrs ...any)
	Warn(msg string, attrs ...any)
	Error(msg string, attrs ...any)
	LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr)
}

type logger struct {
	logger *slog.Logger
}

type replaceAttr func(groups []string, a slog.Attr) slog.Attr

func InitLogger(cfg *config.Config) *logger {
	logLevel := setLogLevel(cfg.Environment)
	replace := replaceLoggerAttr()
	logger := newLogger(replace, logLevel)

	return logger
}

func newLogger(replace replaceAttr, logLevel slog.Level) *logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: replace,
		Level:       logLevel,
	})

	return &logger{
		logger: slog.New(handler),
	}
}

func replaceLoggerAttr() replaceAttr {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		// Remove the time key if the log is not a group.
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}

		if a.Key == slog.SourceKey {
			return slog.Attr{}
		}

		return a
	}

	return replace
}

func (l *logger) log(level slog.Level, msg string, attrs ...any) {
	// get the caller file and line which is 2 level above the current function
	_, file, line, _ := runtime.Caller(2)

	// include the source file and line in the log
	group := slog.Group(
		"source",
		slog.Attr{
			Key:   "filename",
			Value: slog.AnyValue(file),
		},
		slog.Attr{
			Key:   "line",
			Value: slog.AnyValue(line),
		},
	)
	attrs = append(attrs, group)
	l.logger.Log(context.Background(), level, msg, attrs...)
}

func setLogLevel(env string) slog.Level {
	if env == "local" {
		return slog.LevelDebug
	}

	return slog.LevelInfo
}

func (l *logger) Debug(msg string, attrs ...any) {
	l.log(slog.LevelDebug, msg, attrs...)
}

func (l *logger) Info(msg string, attrs ...any) {
	l.log(slog.LevelInfo, msg, attrs...)
}

func (l *logger) Warn(msg string, attrs ...any) {
	l.log(slog.LevelWarn, msg, attrs...)
}

func (l *logger) Error(msg string, attrs ...any) {
	l.log(slog.LevelError, msg, attrs...)
}

func (l *logger) LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	l.logger.LogAttrs(ctx, level, msg, attrs...)
}
