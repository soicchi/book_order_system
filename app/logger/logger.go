package logger

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/soicchi/book_order_system/config"
)

type replaceAttr func(groups []string, a slog.Attr) slog.Attr

func InitLogger(cfg config.Config) *slog.Logger {
	logLevel := setLogLevel(cfg.Env)
	replace := replaceLoggerAttr()
	logger := newLogger(replace, logLevel)

	return logger
}

func newLogger(replace replaceAttr, logLevel slog.Level) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: replace,
		Level:       logLevel,
	})

	return slog.New(handler)
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
