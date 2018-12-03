package main

import (
	"email-bot/emailbot"
	"email-bot/helpers/generalhelpers"
	"email-bot/lg"
	"fmt"
)

func main() {
	bm := emailbot.NewManager()
	p := emailbot.NewProfileManager().NewExistingProfile()
	p.Populate()
	fmt.Println(p.Profiles())

	bm.SetProfile(p)

	loginAddress := "https://passport.yandex.com/auth"
	loginInput := "//input[@name='login']"
	passwordInput := "//input[@name='passwd']"
	signIn := "//button[@class='passport-Button']"

	inboxAddress := "https://mail.yandex.com/#compose"
	toField := "//div[@name='to']"
	subjectField := "//input[contains(@name, 'subj-')]"
	contentField := "//div[@class='cke_wysiwyg_div cke_reset cke_enable_context_menu cke_editable cke_editable_themed cke_contents_ltr cke_show_borders']"

	sendButton := "//button[contains(@class, 'js-send nb-group-start')]"

	bm.AddAction(false).AddVisit(loginAddress).AddWait(300)

	bm.AddAction(false).
		AddFillOperation(loginInput, "email").
		AddFillOperation(passwordInput, "password").
		AddWait(300).
		AddSubmitOperation(signIn).
		AddWait(3000)

	bm.AddAction(false).AddVisit(inboxAddress).AddWait(300)

	bm.AddAction(false).
		AddFillOperation(toField, "idof@live.com.au").
		AddFillOperation(subjectField, "pwnd").
		AddFillOperation(contentField, "etsy getsy").
		AddWait(300).
		AddSubmitOperation(sendButton)

	bm.AddBots(8081, 1)

	bm.ScrapeAll()

	generalhelpers.Wait(200000)
	lg.Debug("end of program")
}
