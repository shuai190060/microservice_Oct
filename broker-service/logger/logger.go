package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger(logFile string) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Errorf("failed to open the log file %s for saving the logs:%v", logFile, err)

	} else {
		Log.Out = file
	}
	Log.SetLevel(logrus.InfoLevel)
}
