package master

import (
	"email-bot/database"
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
	archivist := database.NewArchivist()
	archivist.RecordProfile(profile)
	archivist.Close()
}

func (m *Master) generatedData() map[string][]string {
	profile := m.profileManager.StandardProfile()
	profile.Generate()
	return profile.Values
}

func (m *Master) processResults(success bool) {
	if success {
		m.saveProfile(m.scrapeManager.ActiveProfileData())
	}
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (m *Master) Scrape() {
	m.scrapeManager.AddValues(m.generatedData())
	m.scrapeManager.ProvisionHotmailNewAccountScrape()
	m.processResults(m.scrapeManager.Scrape())
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
