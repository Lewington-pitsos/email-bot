package master

import (
	"email-bot/offline/files"
	"email-bot/offline/profile"
	"email-bot/online/scrape"
	"go/build"
)

var profileBankPath = build.Default.GOPATH + "src/email-bot/data/"
var profileList = build.Default.GOPATH + "src/email-bot/data/profiles.json"

// +---------------------------------------------------------------------------------------+
//										Master STRUCT
// +---------------------------------------------------------------------------------------+

type Master struct {
	profileManager *profile.Manager
	scrapeManager  *scrape.Manager
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (m *Master) saveProfile(profile map[string]string) {
	saver := files.NewManager()
	saver.RecordProfile(profile["email"], profile)
}

func (m *Master) generatedData() map[string][]string {
	profile := m.profileManager.StandardProfile()
	profile.Generate()
	return profile.Values
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (m *Master) Scrape() {
	m.scrapeManager.AddValues(m.generatedData())
	m.scrapeManager.ProvisionHotmailNewAccountScrape()
	m.scrapeManager.Scrape()
	m.saveProfile(m.scrapeManager.ActiveProfileData())
}

// +---------------------------------------------------------------------------------------+
//									EXPOSED FUNCTIONS
// +---------------------------------------------------------------------------------------+

func NewMaster() *Master {
	return &Master{
		profileManager: profile.NewManager(),
		scrapeManager:  scrape.NewManager(8081),
	}
}
