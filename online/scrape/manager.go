package scrape

import (
	"email-bot/datastructures"
	"email-bot/lg"
	"email-bot/online/action"
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

func (m *Manager) ActiveProfileData() map[string]string {
	profile := make(map[string]string)
	lg.Debug("Extracting data for entered profile")

	for name, detail := range m.candidateValues {
		profile[name] = detail.CurrentValue()
	}

	profile["birthdate"] = m.birthDate(profile)
	lg.Debug(profile["birthdate"])

	return profile
}

func (m *Manager) AddValues(values map[string]datastructures.Detail) {
	m.candidateValues = values
}

func (m *Manager) ProvisionScrape(actions []*action.Action) {
	for _, action := range actions {
		action.AddCandidateValues(m.candidateValues)
		action.AddBrowser(m.browser)
		m.scrape.AddAction(action)
	}
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
