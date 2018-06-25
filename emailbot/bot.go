package emailbot

import (
	"email-bot/database"
	"email-bot/datastructures"
	"email-bot/offline/profile"
	"email-bot/online/scrape"
	"go/build"
)

var profileBankPath = build.Default.GOPATH + "src/email-bot/data/"
var profileList = build.Default.GOPATH + "src/email-bot/data/profiles.json"

// +---------------------------------------------------------------------------------------+
//										Bot STRUCT
// +---------------------------------------------------------------------------------------+

type Bot struct {
	dataProfile   *profile.DataProfile
	scrapeManager *scrape.Manager
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (m *Bot) saveProfile(profile map[string]string) {
	archivist := database.NewArchivist()
	archivist.RecordProfile(profile)
	archivist.Close()
}

func (m *Bot) generatedData() map[string]datastructures.Detail {
	profile.Generate()
	return profile.Values
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
	m.scrapeManager.AddValues(m.generatedData())
	m.scrapeManager.ProvisionHotmailNewAccountScrape()
	m.processResults(m.scrapeManager.Scrape())
}

// +---------------------------------------------------------------------------------------+
//									EXPOSED FUNCTIONS
// +---------------------------------------------------------------------------------------+

func NewBot(port int, profile *profile.DataProfile) *Bot {
	return &Bot{
		dataProfile:   profile,
		scrapeManager: scrape.NewManager(port),
	}
}
