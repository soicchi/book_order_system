package logging

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/soicchi/book_order_system/internal/config"
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
			source := a.Value.Any().(*slog.Source)

			// Remove the directory from the source's filename.
			source.File = filepath.Base(source.File)

			// Display only the function name
			splitFunctionName := strings.Split(source.Function, "/")
			source.Function = splitFunctionName[len(splitFunctionName)-1]
		}

		return a
	}

	return replace
}

func setLogLevel(env string) slog.Level {
	if env == "local" {
		return slog.LevelDebug
	}

	return slog.LevelInfo
}

func (l *logger) Debug(msg string, attrs ...any) {
	l.logger.Debug(msg, attrs...)
}

func (l *logger) Info(msg string, attrs ...any) {
	l.logger.Info(msg, attrs...)
}

func (l *logger) Warn(msg string, attrs ...any) {
	l.logger.Warn(msg, attrs...)
}

func (l *logger) Error(msg string, attrs ...any) {
	l.logger.Error(msg, attrs...)
}

func (l *logger) LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	l.logger.LogAttrs(ctx, level, msg, attrs...)
}
