package logger

import (
	"go/build"
	"log"
	"os"
)


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

var LoggerInterface *log.Logger

func init() {
	LoggerInterface = actualLogger("scrape")
}