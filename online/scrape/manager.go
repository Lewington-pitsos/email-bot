package scrape

import (
	"email-bot/logger"
	"email-bot/online/browser"
	"email-bot/online/data"
	"fmt"
)

// +---------------------------------------------------------------------------------------+
//										Manager STRUCT
// +---------------------------------------------------------------------------------------+

type Manager struct {
	candidateValues map[string]*data.Detail
	scrape          *Scrape
	browser         *browser.Browser
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (m *Manager) birthDate(profile map[string]string) string {
	return fmt.Sprintf("%s %s %s", profile["day"], profile["month"], profile["year"])
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (m *Manager) Scrape() bool {
	m.scrape.Scrape()
	m.browser.Quit()
	return m.scrape.Success
}

func (m *Manager) ActiveProfileData() map[string]string {
	profile := make(map[string]string)
	logger.LoggerInterface.Println("Extracting data for entered profile")

	for name, detail := range m.candidateValues {
		profile[name] = detail.CurrentValue()
	}

	profile["birthdate"] = m.birthDate(profile)

	return profile
}

func (m *Manager) AddValues(values map[string][]string) {
	for name, valueSlice := range values {
		m.candidateValues[name] = data.NewDetail(valueSlice)
	}
}

func (m *Manager) ProvisionHotmailNewAccountScrape() {
	actions := hotmailNewAccountScrapeActions(m.candidateValues, m.browser)
	m.scrape.AddActions(actions)
}

// +---------------------------------------------------------------------------------------+
//										Manager STRUCT
// +---------------------------------------------------------------------------------------+

func NewManager(port int) *Manager {
	return &Manager{
		candidateValues: make(map[string]*data.Detail),
		browser:         browser.NewBrowser(port),
		scrape:          NewScrape(),
	}
}
