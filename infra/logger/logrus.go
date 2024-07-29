package logger

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger() Logger {
	l := logrus.New()
	// Configurações adicionais, como formato, nível de log, etc.
	return &LogrusLogger{log: l}
}

func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *LogrusLogger) Error(format string) {
	l.log.Error(format)

}
func (l *LogrusLogger) Fatal(err interface{}) {
	l.log.Fatal(err)

}
func (l *LogrusLogger) Fatalf(format string, err interface{}) {
	l.log.Fatalf(format, err)

}
