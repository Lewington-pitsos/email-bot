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
	save          bool
	profile       map[string]datastructures.Detail
	scrapeManager *scrape.Manager
	actions       []*action.Action
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (b *Bot) saveProfile(profile map[string]string) {
	archivist := database.NewArchivist()
	archivist.RecordYandexProfile(profile)
	archivist.Close()
}

func (b *Bot) processResults(success bool) {
	if success && b.save {
		b.saveProfile(b.scrapeManager.ActiveProfileData())
	}
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (b *Bot) Scrape() {
	b.scrapeManager.AddValues(b.profile)
	b.scrapeManager.ProvisionScrape(b.actions)
	b.processResults(b.scrapeManager.Scrape())
}

// +---------------------------------------------------------------------------------------+
//									EXPOSED FUNCTIONS
// +---------------------------------------------------------------------------------------+

func NewBot(port int, profile map[string]datastructures.Detail, actions []*action.Action, save bool) *Bot {
	return &Bot{
		save:          save,
		profile:       profile,
		scrapeManager: scrape.NewManager(port),
		actions:       actions,
	}
}
