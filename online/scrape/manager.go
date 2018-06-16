package scrape

import (
	"email-bot/online/action"
	"email-bot/online/crawler"

	"github.com/tebeka/selenium"
)

type Manager struct {
	values  map[string][]string
	scrape  *Scrape
	browser selenium.WebDriver
}

func (m *Manager) AddAction(action *action.Action) {
	m.scrape.AddActions(action)
}

func (m *Manager) Scrape() {
	m.scrape.Scrape()
}

func NewManager(port int, values map[string][]string) *Manager {
	return &Manager{
		values:  values,
		browser: crawler.NewWebDriver(port),
		scrape:  NewScrape(),
	}
}

func (m *Manager) ProvisionHotmailNewAccount() {
	emailInput := "//input[@id='MemberName']"
	passInput := "//input[@id='PasswordInput']"
	firstInput := "//input[@id='FirstName']"
	lastInput := "//input[@id='LastName']"
	dayInput := "//select[@id='BirthDay']"
	monthInput := "//select[@id='BirthMonth']"
	yearInput := "//select[@id='BirthYear']"

	submitInput := "//input[@id='iSignupAction']"

	a := action.NewAction(m.browser)
	command := action.VisitPage("https://signup.live.com/signup")
	a.AddToInteraction(command)

	// ========================================================

	a2 := action.NewAction(m.browser)
	a2.AddFillOperation(emailInput, m.values["email"][0])
	a2.AddSubmitOperation(submitInput)

	// ========================================================

	a3 := action.NewAction(m.browser)
	a3.AddFillOperation(passInput, "mj8fdsayum9o1ws")
	a3.AddSubmitOperation(submitInput)

	// ========================================================

	a4 := action.NewAction(m.browser)
	a4.AddFillOperation(firstInput, m.values["username"][0])
	a4.AddFillOperation(lastInput, m.values["username"][1])
	a4.AddSubmitOperation(submitInput)

	// ========================================================

	a5 := action.NewAction(m.browser)
	a5.AddFillOperation(dayInput, "2")
	a5.AddFillOperation(monthInput, "January")
	a5.AddFillOperation(yearInput, "1992")
	a5.AddToSpec(action.CheckExists(submitInput))
	a5.AddToInteraction(action.Click(submitInput))

	// ========================================================

	m.scrape.AddActions(a, a2, a3, a4, a5)
}
