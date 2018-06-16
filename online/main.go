package online

import (
	"email-bot/online/action"
	"email-bot/online/crawler"
	"email-bot/online/scrape"
	"fmt"
	"time"
)

func Test() {
	emailInput := "//input[@id='MemberName']"
	passInput := "//input[@id='PasswordInput']"
	firstInput := "//input[@id='FirstName']"
	lastInput := "//input[@id='LastName']"
	dayInput := "//select[@id='BirthDay']"
	monthInput := "//select[@id='BirthMonth']"
	yearInput := "//select[@id='BirthYear']"

	submitInput := "//input[@id='iSignupAction']"

	wd := crawler.NewWebDriver(8081)
	a := action.NewAction(wd)
	command := action.VisitPage("https://signup.live.com/signup")
	a.AddToInteraction(command)

	// ========================================================

	a2 := action.NewAction(wd)
	spec1 := action.CheckExists(emailInput)
	spec2 := action.CheckExists(submitInput)
	command2 := action.FillField(emailInput, "jacobKerrim@hotmail.com")
	command3 := action.Click(submitInput)

	a2.AddToSpec(spec1).AddToSpec(spec2).AddToInteraction(command2).AddToInteraction(command3)

	// ========================================================

	a3 := action.NewAction(wd)
	spec3 := action.CheckExists(passInput)
	spec4 := action.CheckExists(submitInput)
	command4 := action.FillField(passInput, "nhsda63218jg7t8")
	command5 := action.Click(submitInput)

	a3.AddToSpec(spec3).AddToSpec(spec4).AddToInteraction(command4).AddToInteraction(command5)

	// ========================================================

	a4 := action.NewAction(wd)
	a4.AddToSpec(action.CheckExists(firstInput))
	a4.AddToSpec(action.CheckExists(lastInput))
	a4.AddToSpec(action.CheckExists(submitInput))
	a4.AddToInteraction(action.FillField(firstInput, "frank"))
	a4.AddToInteraction(action.FillField(lastInput, "mithen"))
	a4.AddToInteraction(action.Click(submitInput))

	// ========================================================

	a5 := action.NewAction(wd)
	a5.AddFillOperation(dayInput, "2")
	a5.AddFillOperation(monthInput, "January")
	a5.AddFillOperation(yearInput, "1992")
	a5.AddToSpec(action.CheckExists(submitInput))
	a5.AddToInteraction(action.Click(submitInput))

	// ========================================================

	scrape := scrape.NewScrape()
	scrape.AddAction(a)
	scrape.AddAction(a2)
	scrape.AddAction(a3)
	scrape.AddAction(a4)
	scrape.AddAction(a5)

	fmt.Println(scrape)
	scrape.Scrape()

	time.Sleep(time.Millisecond * 80000)
}
