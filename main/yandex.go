package main

import (
	"email-bot/emailbot"
	"email-bot/helpers/generalhelpers"
	"email-bot/logger"
	"email-bot/online/action"
)

func main() {
	botManager := emailbot.NewManager()
	p := emailbot.NewProfileManager().NewDataProfile()

	p.AddToBank("name", "female-en").
		AddToBank("slang", "internet-slang").
		AddToBank("datevault-day", "2").
		AddToBank("datevault-month", "1").
		AddToBank("datevault-year", "2006").
		AddToBank("passvault", "")

	p.AddField("firstname", 1).
		AddField("firstname", 1).
		WithChunk("bank", "name").
		AddField("lastname", 1).
		WithChunk("bank", "name").
		AddField("fullname", 1).
		WithChunk("derived", "firstname").
		WithChunk("derived", "lastname").
		AddField("email", 30).
		WithModifiedChunk("derived", "fullname", "slang").
		AddField("password", 1).
		WithChunk("bank", "passvault").
		AddField("question", 1).
		WithChunk("literal", "Your favorite actor or actress").
		AddField("answer", 1).
		WithChunk("bank", "name").
		WithChunk("literal", " ").
		WithChunk("bank", "name")

	botManager.SetProfile(p)

	firstInput := "//input[@id='firstname']"
	lastInput := "//input[@id='lastname']"
	emailInput := "//input[@id='login']"
	passInput := "//input[@id='password']"
	passConfirm := "//input[@id='password_confirm']"
	questionSelector := "//button[contains(@class, 'control-questions')]"
	answerSelector := "//span[contains(text(), '%s')]"
	answerInput := "//input[@id='hint_answer']"
	errorMessage := "//*[@class='error-message']"

	noPhoneButton := "//span[contains(@class, 'link_has-no-phone')]"
	//submitButton := "//button[contains(@class, 'button2_type_submit js-submit')]"
	avatar := "//span[@class='avatar-mask']"
	copywright := "//div[@class='n-footer__rights']"

	botManager.AddAction(false).
		AddVisit("https://passport.yandex.com/registration").
		AddWait(300)

	// ========================================================

	botManager.AddAction(false).
		AddClick(firstInput).
		AddFillOperation(firstInput, "firstname").
		AddClick(lastInput).
		AddFillOperation(lastInput, "lastname").
		AddFillOperation(emailInput, "email").
		AddFillOperation(passInput, "password").
		AddFillOperation(passConfirm, "password").
		AddSubmitOperation(noPhoneButton).
		AddWait(300)

		// ========================================================

	botManager.AddAction(false).
		AddSelectOperation(questionSelector, answerSelector, "question").
		AddWait(400).
		AddFillOperation(answerInput, "answer").
		AddClick(lastInput).
		AddWait(4000)

		// ========================================================

	botManager.AddAction(false).
		AddToSpec(action.CheckDoesntExist(errorMessage))

	// ========================================================

	botManager.AddAction(true).
		AddContinuousCheck([]string{avatar, copywright},
			2000,
			200,
		)

	botManager.AddBots(8081, 1)

	botManager.ScrapeAll()

	generalhelpers.Wait(200000)
	logger.LoggerInterface.Println("end of program")
}
