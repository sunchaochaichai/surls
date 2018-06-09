package global

import (
	"github.com/sirupsen/logrus"
	"os"
)

func newLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		DisableColors:    false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05.0000",
		QuoteEmptyFields: true,
	}
	logFile, err := os.OpenFile(
		LogPath+"/app.log",
		os.O_CREATE|os.O_WRONLY,
		os.ModePerm,
	)

	if err != nil {
		logrus.Fatal("log file create failed.", err)
	}

	logger.Out = logFile

	return logger
}
