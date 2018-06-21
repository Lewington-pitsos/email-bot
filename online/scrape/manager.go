package scrape

import (
	"email-bot/online/action"
	"email-bot/online/browser"
	"email-bot/online/data"
)

type Manager struct {
	details map[string]*data.Detail
	scrape  *Scrape
	browser *browser.Browser
}

func (m *Manager) AddAction(action *action.Action) {
	m.scrape.AddActions(action)
}

func (m *Manager) Scrape() {
	m.scrape.Scrape()
	m.browser.Quit()
}

func (m *Manager) ProfileData() map[string]string {
	profile := make(map[string]string)

	for name, detail := range m.details {
		profile[name] = detail.CurrentValue()
	}

	return profile
}

func (m *Manager) AddValues(values map[string][]string) {
	for name, valueSlice := range values {
		m.details[name] = data.NewDetail(valueSlice)
	}
}

func NewManager(port int) *Manager {
	return &Manager{
		details: make(map[string]*data.Detail),
		browser: browser.NewBrowser(port),
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
	capchaBox := "//div[@id='hipTemplateContainer']"

	submitInput := "//input[@id='iSignupAction']"

	a := action.NewAction(m.browser)
	command := action.VisitPage("https://signup.live.com/signup")
	a.AddToInteraction(command)

	// ========================================================

	a2 := action.NewAction(m.browser)
	a2.AddFillOperation(emailInput, m.details["email"])
	a2.AddSubmitOperation(submitInput)

	// ========================================================

	a3 := action.NewAction(m.browser)
	a3.AddFillOperation(passInput, m.details["password"])
	a3.AddSubmitOperation(submitInput)

	// ========================================================

	a4 := action.NewAction(m.browser)
	a4.AddFillOperation(firstInput, m.details["username"])
	a4.AddFillOperation(lastInput, m.details["username"])
	a4.AddSubmitOperation(submitInput)

	// ========================================================

	a5 := action.NewAction(m.browser)
	a5.AddFillOperation(dayInput, m.details["day"])
	a5.AddFillOperation(monthInput, m.details["month"])
	a5.AddFillOperation(yearInput, m.details["year"])
	a5.AddToSpec(action.CheckExists(submitInput))
	a5.AddToInteraction(action.Click(submitInput))

	// ========================================================

	a6 := action.NewAction(m.browser)
	a6.AddToSpec(action.CheckExists(capchaBox))

	m.scrape.AddActions(a, a2, a3, a4, a5, a6)
}
