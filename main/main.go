package main

import (
	"email-bot/logger"
	"email-bot/offline"
)

func main() {
	logger.Log.Println("and awaaaaaaay we go")
	offline.Test()
}
