package main

import (
	"email-bot/emailbot"
	"email-bot/helpers/generalhelpers"
	"email-bot/logger"
	"fmt"
)

func main() {
	bm := emailbot.NewManager()
	p := emailbot.NewProfileManager().NewExistingProfile()
	p.Populate()
	fmt.Println(p.Profiles())

	bm.SetProfile(p)

	address := "https://passport.yandex.com/auth"
	loginInput := "//input[@name='login']"
	passwordInput := "//input[@name='passwd']"
	signIn := "//button[@class='passport-Button']"

	bm.AddAction(false).AddVisit(address).AddWait(300)

	bm.AddAction(false).
		AddFillOperation(loginInput, "email").
		AddFillOperation(passwordInput, "password").
		AddWait(300).
		AddSubmitOperation(signIn).
		AddWait(90000)

	bm.AddBots(8081, 1)

	bm.ScrapeAll()

	generalhelpers.Wait(200000)
	logger.LoggerInterface.Println("end of program")
}
