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

	submitInput := "//input[@id='iSignupAction']"

	a := action.NewAction(browser)
	command := action.VisitPage("https://signup.live.com/signup")
	a.AddToInteraction(command)
	a.AddWait(3000)

	// ========================================================

	a2 := action.NewAction(browser)
	a2.AddFillOperation(emailInput, candidateValues["email"])
	a2.AddWait(200)
	a2.AddSubmitOperation(submitInput)
	a2.AddWait(300)

	// ========================================================

	a3 := action.NewAction(browser)
	a3.AddFillOperation(passInput, candidateValues["password"])
	a3.AddWait(200)
	a3.AddSubmitOperation(submitInput)
	a3.AddWait(300)

	// ========================================================

	a4 := action.NewAction(browser)
	a4.AddFillOperation(firstInput, candidateValues["username"])
	a4.AddFillOperation(lastInput, candidateValues["username"])
	a4.AddWait(200)
	a4.AddSubmitOperation(submitInput)
	a4.AddWait(300)

	// ========================================================

	a5 := action.NewAction(browser)
	a5.AddFillOperation(dayInput, candidateValues["day"])
	a5.AddFillOperation(monthInput, candidateValues["month"])
	a5.AddFillOperation(yearInput, candidateValues["year"])
	a5.AddWait(300)
	a5.AddToSpec(action.CheckExists(submitInput))
	a5.AddToInteraction(action.Click(submitInput))
	a5.AddWait(3000)

	// ========================================================

	a6 := action.NewAction(browser)
	a6.AddToSpec(action.CheckExists(capchaBox))
	a6.AddWait(100000)

	// ========================================================

	a7 := action.NewAction(browser)
	a7.AddToSpec(action.CheckExists(capchaBox))
	a6.AddWait(100000)

	return []*action.Action{a, a2, a3, a4, a5, a6, a7}
}
