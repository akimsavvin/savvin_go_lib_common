package logger

import (
	"log/slog"
	"os"

	"github.com/akimsavvin/savvin_go_lib_common/config"
	"github.com/akimsavvin/savvin_go_lib_common/logger/sl"
)

func Init(env string) *slog.Logger {
	log := new(slog.Logger)

	switch env {
	case config.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case config.EnvTest:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}

	log.Debug(
		"logger initialized",
		sl.Pkg("logger"),
		sl.Op("Init"),
	)

	return log
}
