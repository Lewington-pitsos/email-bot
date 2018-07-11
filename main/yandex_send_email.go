package main

import (
	"email-bot/emailbot"
)

func main() {
	bm := emailbot.NewManager()
	p := emailbot.NewProfileManager().NewExistingProfile()
}
