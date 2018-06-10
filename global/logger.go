package global

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Logger
	Fields logrus.Fields
}

func (this Logger) Log(keyvals ...interface{}) error {
	this.Logger.WithFields(this.Fields).Info(keyvals)
	return nil
}

func newLogger() Logger {

	l := Logger{
		logrus.New(),
		make(logrus.Fields),
	}
	l.Formatter = &logrus.TextFormatter{
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

	l.Out = logFile

	return l
}
