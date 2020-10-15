package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	infoLogger  log.Logger
	warnLogger  log.Logger
	errorLogger log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:  *log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		warnLogger:  *log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime),
		errorLogger: *log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime),
	}
}

func (l *Logger) Info(msg interface{}) {
	l.infoLogger.Println(msg)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Println(fmt.Sprintf(format, v))
}

func (l *Logger) Warn(msg interface{}) {
	l.warnLogger.Println(msg)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warnLogger.Println(fmt.Sprintf(format, v))
}

func (l *Logger) Error(msg interface{}) {
	l.errorLogger.Println(msg)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Println(fmt.Sprintf(format, v))
}
