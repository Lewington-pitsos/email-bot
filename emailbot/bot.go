package emailbot

import (
	"email-bot/online/scrape"
)

// +---------------------------------------------------------------------------------------+
//										Bot STRUCT
// +---------------------------------------------------------------------------------------+

type Bot struct {
	scrapeManager *scrape.Manager
	scrape.Profile
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (b *Bot) Scrape() {
	b.scrapeManager.ProvisionScrape(b.Instructions())
	b.scrapeManager.Scrape()
}

// +---------------------------------------------------------------------------------------+
//									EXPOSED FUNCTIONS
// +---------------------------------------------------------------------------------------+

func NewBot(port int) *Bot {
	return &Bot{
		scrapeManager: scrape.NewManager(port),
	}
}
