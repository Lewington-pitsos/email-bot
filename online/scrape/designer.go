package scrape

import (
	"email-bot/online/action"
	"email-bot/online/browser"
	"email-bot/online/data"
)

func hotmailNewAccountScrapeActions(candidateValues map[string]*data.Detail, browser *browser.Browser) []*action.Action {
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

	a := action.NewAction(browser, false)
	command := action.VisitPage("https://signup.live.com/signup")
	a.AddToInteraction(command)
	a.AddWait(3000)

	// ========================================================

	a2 := action.NewAction(browser, false)
	a2.AddFillOperation(emailInput, candidateValues["email"])
	a2.AddWait(200)
	a2.AddSubmitOperation(submitInput)
	a2.AddWait(300)

	// ========================================================

	a3 := action.NewAction(browser, false)
	a3.AddFillOperation(passInput, candidateValues["password"])
	a3.AddWait(200)
	a3.AddSubmitOperation(submitInput)
	a3.AddWait(300)

	// ========================================================

	a4 := action.NewAction(browser, false)
	a4.AddFillOperation(firstInput, candidateValues["firstname"])
	a4.AddFillOperation(lastInput, candidateValues["lastname"])
	a4.AddWait(200)
	a4.AddSubmitOperation(submitInput)
	a4.AddWait(300)

	// ========================================================

	a5 := action.NewAction(browser, false)
	a5.AddFillOperation(dayInput, candidateValues["day"])
	a5.AddFillOperation(monthInput, candidateValues["month"])
	a5.AddFillOperation(yearInput, candidateValues["year"])
	a5.AddWait(300)
	a5.AddToSpec(action.CheckExists(submitInput))
	a5.AddToInteraction(action.Click(submitInput))
	a5.AddWait(4000)

	// ========================================================

	a6 := action.NewAction(browser, false)
	a6.AddToSpec(action.CheckExists(capchaBox))
	a6.AddWait(10000)

	// ========================================================

	a7 := action.NewAction(browser, true)
	a7.AddToSpec(action.CheckExists(homeBanner))
	a7.AddToSpec(action.CheckExists(homeHeader))

	return []*action.Action{a, a2, a3, a4, a5, a6, a7}
}
