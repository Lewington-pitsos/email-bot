package lg

import (
	"fmt"
	"go/build"
	"os"

	log "github.com/sirupsen/logrus"
)

var input = make(chan statement, 200)

func keepLogging(statementQueue <-chan statement) {

	for s := range statementQueue {
		switch s.kind {
		case "DEBUG":
			log.Debug(s.message)
		case "INFO":
			log.Info(s.message)
		case "WARN":
			log.Warn(s.message)
		case "ERROR":
			log.Error(s.message)
		}
	}
}

func init() {

	var logpath = build.Default.GOPATH + "/src/email-bot/lg/main.log"

	var file, err1 = os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE, 0755)

	if err1 != nil {
		panic(err1)
	}

	log.SetLevel(log.DebugLevel)
	log.SetOutput(file)

	go keepLogging(input)
}

func Debug(str string, args ...interface{}) {
	input <- statement{
		kind:    "DEBUG",
		message: fmt.Sprintf(str, args...),
	}
}

func Info(str string, args ...interface{}) {
	input <- statement{
		kind:    "INFO",
		message: fmt.Sprintf(str, args...),
	}
}

func Warn(str string, args ...interface{}) {
	input <- statement{
		kind:    "WARN",
		message: fmt.Sprintf(str, args...),
	}
}

func Error(str string, args ...interface{}) {
	input <- statement{
		kind:    "ERROR",
		message: fmt.Sprintf(str, args...),
	}
}

func DebugMap(preface string, stringMap map[string]string) {
	Debug(preface)
	for key, value := range stringMap {
		Debug("%s: %v", key, value)
	}
}

func DebugMapSlice(preface string, sliceMap map[string][]string) {
	Debug(preface)
	for key, slice := range sliceMap {
		Debug("%s: %v", key, slice)
	}
}
