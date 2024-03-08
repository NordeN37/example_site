package logger

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

func New(levelStr string) *zerolog.Logger {
	var level zerolog.Level
	switch levelStr {
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "error":
		level = zerolog.ErrorLevel
	default:
		level = zerolog.DebugLevel
	}
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().Local()
	}
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		count := 2
		for i := len(file) - 1; i > 0; i-- {
			if count == 0 {
				break
			}
			if file[i] == '/' {
				count -= 1
				short = file[i+1:]
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
	log := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger().Level(level)
	return &log
}
