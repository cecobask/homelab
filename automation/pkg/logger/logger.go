package logger

import (
	"io"
	"log/slog"
)

func NewLogger(writer io.Writer) *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(writer, opts)
	return slog.New(handler)
}
