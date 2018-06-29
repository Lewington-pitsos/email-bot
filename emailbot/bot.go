package emailbot

import (
	"email-bot/database"
	"email-bot/datastructures"
	"email-bot/logger"
	"email-bot/offline/profile"
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
	profile       profile.ProfileInterface
	scrapeManager *scrape.Manager
	actions       []*action.Action
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (b *Bot) saveProfile(profile map[string]string) {
	archivist := database.NewArchivist()
	archivist.RecordProfile(profile)
	archivist.Close()
}

func (b *Bot) processResults(success bool) {
	if success {
		b.saveProfile(b.scrapeManager.ActiveProfileData())
	}
}

func (b *Bot) generatedValues() map[string]datastructures.Detail {
	b.profile.Generate()
	return b.profile.Values()
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (b *Bot) Scrape() {
	logger.LoggerInterface.Println(b.generatedValues())
	b.scrapeManager.AddValues(b.generatedValues())
	b.scrapeManager.ProvisionScrape(b.actions)
	b.processResults(b.scrapeManager.Scrape())
}

// +---------------------------------------------------------------------------------------+
//									EXPOSED FUNCTIONS
// +---------------------------------------------------------------------------------------+

func NewBot(port int, profile profile.ProfileInterface, actions []*action.Action) *Bot {
	return &Bot{
		profile:       profile,
		scrapeManager: scrape.NewManager(port),
		actions:       actions,
	}
}
