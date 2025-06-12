package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func InitLogger() {
	handler := slog.NewTextHandler(os.Stdout, nil)

	Log = slog.New(handler)
}
