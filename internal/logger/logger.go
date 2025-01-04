package logger

import (
	"io"
	"log/slog"
	"os"
	"runtime"
)

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Error(err error, args ...any) {
	slog.Error(err.Error(), args...)
}

func Init() {
	var lout io.Writer = os.Stdout
	logFile, err := os.OpenFile("./logs/log.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err == nil {
		lout = io.MultiWriter(logFile, os.Stdout)
	}

	skipDepth := 8 //found experimentally
	logger := slog.New(slog.NewJSONHandler(
		lout,
		&slog.HandlerOptions{
			AddSource: true,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.SourceKey {
					pc := make([]uintptr, 1)
					l := runtime.Callers(skipDepth, pc[:])
					frames := runtime.CallersFrames(pc[:l])
					frame, _ := frames.Next()

					a.Value = slog.GroupValue(
						slog.String("function", frame.Function),
						slog.String("file", frame.File),
						slog.Int("line", frame.Line),
					)
				}

				return a
			},
		}),
	)

	slog.SetDefault(logger)
}
