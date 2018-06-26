package emailbot

import (
	"email-bot/database"
	"email-bot/datastructures"
	"email-bot/online/action"
	"email-bot/online/scrape"
	"go/build"
)

var profileBankPath = build.Default.GOPATH + "src/email-bot/data/"
var profileList = build.Default.GOPATH + "src/email-bot/data/profiles.json"

// +---------------------------------------------------------------------------------------+
//										Bot STRUCT
// +---------------------------------------------------------------------------------------+

type Bot struct {
	data          map[string]datastructures.Detail
	scrapeManager *scrape.Manager
	actions       []*action.Action
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (m *Bot) saveProfile(profile map[string]string) {
	archivist := database.NewArchivist()
	archivist.RecordProfile(profile)
	archivist.Close()
}

func (m *Bot) processResults(success bool) {
	if success {
		m.saveProfile(m.scrapeManager.ActiveProfileData())
	}
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (m *Bot) Scrape() {
	m.scrapeManager.AddValues(m.data)
	m.scrapeManager.ProvisionScrape(m.actions)
	m.processResults(m.scrapeManager.Scrape())
}

// +---------------------------------------------------------------------------------------+
//									EXPOSED FUNCTIONS
// +---------------------------------------------------------------------------------------+

func NewBot(port int, data map[string]datastructures.Detail, actions []*action.Action) *Bot {
	return &Bot{
		data:          data,
		scrapeManager: scrape.NewManager(port),
		actions:       actions,
	}
}
