package scrape

import (
	"email-bot/datastructures"
	"email-bot/logger"
	"email-bot/online/browser"
	"fmt"
)

// +---------------------------------------------------------------------------------------+
//										Manager STRUCT
// +---------------------------------------------------------------------------------------+

type Manager struct {
	candidateValues map[string]datastructures.Detail
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

func(m *Manager) AddAction() {
	action := action.NewAction()
	m.scraoe.AddActions(action)
	return action
}

func (m *Manager) ActiveProfileData() map[string]string {
	profile := make(map[string]string)
	logger.LoggerInterface.Println("Extracting data for entered profile")

	for name, detail := range m.candidateValues {
		profile[name] = detail.CurrentValue()
	}

	profile["birthdate"] = m.birthDate(profile)
	logger.LoggerInterface.Println(profile["birthdate"])

	return profile
}

func (m *Manager) AddValues(values map[string]datastructures.Detail) {
	m.candidateValues = values
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
		candidateValues: make(map[string]datastructures.Detail),
		browser:         browser.NewBrowser(port),
		scrape:          NewScrape(),
	}
}
