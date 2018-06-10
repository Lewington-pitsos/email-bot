package logger

import (
	"go/build"
	"log"
	"os"
)

var (
	// Log the global logger for our project
	Log *log.Logger
)

func init() {
	// set location of log file
	var logpath = build.Default.GOPATH + "/src/email-bot/logger/info.log"

	var file, err1 = os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}
