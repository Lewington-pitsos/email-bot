package logger

import (
	"go/build"
	"log"
	"os"
)

type Logger struct {
	name string
	logger *log.Logger
}

func NewLogger(logName string) *Logger {
	return &Logger{
		name: logName,
		logger: actualLogger(logName),
	}
}

func (l *Logger)Log(message string) {
	l.logger.Println(message)
}

func actualLogger(logName string) *log.Logger {
	// set location of log file
	var logpath = build.Default.GOPATH + "/src/email-bot/logger/" + logName

	var file, err1 = os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err1 != nil {
		panic(err1)
	}

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)
	logger.Println("LogFile : " + logpath)

	return logger;
}
