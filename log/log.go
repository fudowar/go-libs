package log

import (
	"fmt"

	"github.com/fudowar/go-libs/time"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: time.NormalFormatStr,
		FullTimestamp:   true,
	})
}

func Info(format string, a ...any) {
	log.Info(fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	log.Error(fmt.Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	log.Warn(fmt.Sprintf(format, a...))
}
