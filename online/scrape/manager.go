package scrape

import (
	"email-bot/online/browser"
	"email-bot/online/data"
)

type Manager struct {
	candidateValues map[string]*data.Detail
	scrape  *Scrape
	browser *browser.Browser
}

func (m *Manager) Scrape() {
	m.scrape.Scrape()
	m.browser.Quit()
}

func (m *Manager) ProfileData() map[string]string {
	profile := make(map[string]string)

	for name, detail := range m.candidateValues {
		profile[name] = detail.CurrentValue()
	}

	return profile
}

func (m *Manager) AddValues(values map[string][]string) {
	for name, valueSlice := range values {
		m.candidateValues[name] = data.NewDetail(valueSlice)
	}
}

func NewManager(port int) *Manager {
	return &Manager{
		candidateValues: make(map[string]*data.Detail),
		browser: browser.NewBrowser(port),
		scrape:  NewScrape(),
	}
}

func (m *Manager) ProvisionHotmailNewAccount() {
	actions := ProvisionHotmailNewAccount(m.candidateValues)
	m.scrape.AddActions(actions)
}
