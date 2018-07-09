package main

import (
	"email-bot/emailbot"
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
		WithChunk("bank", "name").
		AddField("lastname", 1).
		WithChunk("bank", "name").
		AddField("fullname", 1).
		WithChunk("derived", "firstname").
		WithChunk("derived", "lastname").
		AddField("email", 30).
		WithModifiedChunk("derived", "fullname", "slang").
		WithChunk("literal", "@hotmail.com").
		AddField("day", 1).
		WithChunk("bank", "datevault-day").
		AddField("month", 1).
		WithChunk("bank", "datevault-month").
		AddField("year", 1).
		WithChunk("bank", "datevault-year").
		AddField("password", 20).
		WithChunk("bank", "passvault")

	botManager.SetProfile(p)

	emailInput := "//input[@id='MemberName']"
	passInput := "//input[@id='PasswordInput']"
	firstInput := "//input[@id='FirstName']"
	lastInput := "//input[@id='LastName']"
	dayInput := "//select[@id='BirthDay']"
	monthInput := "//select[@id='BirthMonth']"
	yearInput := "//select[@id='BirthYear']"
	capchaBox := "//div[@id='hipTemplateContainer']"
	homeBanner := "//div[@id='loaded-home-banner-profile-section']"
	homeHeader := "//div[@id='headerUniversalHeader']"

	submitInput := "//input[@id='iSignupAction']"
	dateOption := "//option[@value='%s']"

	botManager.AddAction(false).
		AddVisit("https://signup.live.com/signup").
		AddWait(300)

	// ========================================================

	botManager.AddAction(false).
		AddFillOperation(emailInput, "email").
		AddWait(200).
		AddSubmitOperation(submitInput).
		AddWait(300)

		// ========================================================

	botManager.AddAction(false).
		AddFillOperation(passInput, "password").
		AddWait(200).
		AddSubmitOperation(submitInput).
		AddWait(300)

		// ========================================================

	botManager.AddAction(false).
		AddFillOperation(firstInput, "firstname").
		AddFillOperation(lastInput, "lastname").
		AddWait(200).
		AddSubmitOperation(submitInput).
		AddWait(300)

		// ========================================================

	botManager.AddAction(false).
		AddSelectOperation(dayInput, dateOption, "day").
		AddSelectOperation(monthInput, dateOption, "month").
		AddSelectOperation(yearInput, dateOption, "year").
		AddWait(300).
		AddToSpec(action.CheckExists(submitInput)).
		AddToInteraction(action.Click(submitInput)).
		AddWait(4000)

		// ========================================================

	botManager.AddAction(false).
		AddToSpec(action.CheckExists(capchaBox)).
		AddWait(90000)

	// ========================================================

	botManager.AddAction(true).
		AddToSpec(action.CheckExists(homeBanner)).
		AddToSpec(action.CheckExists(homeHeader))

	botManager.AddBot(8081)

	botManager.ScrapeAll()
}
