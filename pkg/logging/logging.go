package logging

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger = zerolog.Logger

func New() zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05.999",
	}).
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		Caller().
		Logger()
}
