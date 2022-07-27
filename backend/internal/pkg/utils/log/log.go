package log

import (
	"decert/internal/pkg/constants"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log *logrus.Logger

func init() {
	logLevelStr := os.Getenv("LOG_LEVEL")
	if len(logLevelStr) == 0 {
		logLevelStr = "info"
	}
	level, _ := logrus.ParseLevel(logLevelStr)
	log = &logrus.Logger{
		Out: os.Stdout,
		Formatter: &prefixed.TextFormatter{
			FullTimestamp:   true,
			ForceFormatting: true,
			ForceColors:     true,
		},
		Level: level,
	}
}

func Debugf(ctx *gin.Context, format string, params ...interface{}) {
	prefix, _ := ctx.Get(constants.Prefix)
	msg := fmt.Sprintf("["+prefix.(string)+"]"+format, params...)
	log.Debug(msg)
}

func Debugln(ctx *gin.Context, params ...interface{}) {
	prefix, _ := ctx.Get(constants.Prefix)
	msg := "[" + prefix.(string) + "]" + fmt.Sprintln(params...)
	log.Debug(msg)
}

func Errorln(ctx *gin.Context, params ...interface{}) {
	prefix, _ := ctx.Get(constants.Prefix)
	msg := "[" + prefix.(string) + "]" + fmt.Sprintln(params...)
	log.Error(msg)
}

func Errorf(ctx *gin.Context, format string, params ...interface{}) {
	prefix, _ := ctx.Get(constants.Prefix)
	msg := fmt.Sprintf("["+prefix.(string)+"]"+format, params...)
	log.Error(msg)
}
