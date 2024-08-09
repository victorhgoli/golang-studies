package logger

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Error(format string)
	Fatal(err interface{})
	Fatalf(format string, err interface{})
	// Adicione outros métodos conforme necessário (Debug, Warn, etc.)
}
