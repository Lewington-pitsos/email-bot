package main

import (
	"email-bot/emailbot"
	"fmt"
)

func main() {
	//bm := emailbot.NewManager()
	p := emailbot.NewProfileManager().NewExistingProfile()
	p.Populate()
	fmt.Println(p.Profiles())
}
