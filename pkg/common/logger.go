package common

import (
	"log"
	"os"
)

// Logger is used for easier access to different log levels
type Logger struct {
	debugLog   *log.Logger
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
}

var logger Logger

func NewLogger() *Logger {
	return &Logger{
		debugLog:   log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:    log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLog: log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog:   log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l == nil {
		l = NewLogger()
	}

	l.debugLog.Printf(msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if l == nil {
		l = NewLogger()
	}

	l.infoLog.Printf(msg, args...)
}

func (l *Logger) Warning(msg string, args ...interface{}) {
	if l == nil {
		l = NewLogger()
	}

	l.warningLog.Printf(msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	if l == nil {
		l = NewLogger()
	}

	l.errorLog.Printf(msg, args...)
}
