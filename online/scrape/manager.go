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
	a2.AddFillOperation(emailInput, m.values["email"])
	a2.AddSubmitOperation(submitInput)

	// ========================================================

	a3 := action.NewAction(m.browser)
	a3.AddFillOperation(passInput, []string{"mj8fdsayum9o1ws", "mj8fdsayum9-0dfd8s", "mj8fdsayum9-0dfd8s", "mj8fdsayum9-0dfd8s", "mj8fdsayum9-0dfd8s"})
	a3.AddSubmitOperation(submitInput)

	// ========================================================

	a4 := action.NewAction(m.browser)
	a4.AddFillOperation(firstInput, m.values["username"])
	a4.AddFillOperation(lastInput, m.values["username"])
	a4.AddSubmitOperation(submitInput)

	// ========================================================

	a5 := action.NewAction(m.browser)
	a5.AddFillOperation(dayInput, []string{"2", "2", "2", "2"})
	a5.AddFillOperation(monthInput, []string{"January", "January", "January", "January", "January"})
	a5.AddFillOperation(yearInput, []string{"1992", "1992", "1992", "1992", "1992", "1992"})
	a5.AddToSpec(action.CheckExists(submitInput))
	a5.AddToInteraction(action.Click(submitInput))

	// ========================================================

	m.scrape.AddActions(a, a2, a3, a4, a5)
}
