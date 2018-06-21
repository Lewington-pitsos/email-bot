package master

import (
	"email-bot/offline/files"
	"email-bot/offline/profile"
	"email-bot/online/scrape"
	"go/build"
)

var profileBankPath = build.Default.GOPATH + "src/email-bot/data/"
var profileList = build.Default.GOPATH + "src/email-bot/data/profiles.json"

type Master struct {
	profileManager *profile.Manager
	scrapeManager  *scrape.Manager
}

func (m *Master) Scrape() {
	m.scrapeManager.AddValues(m.generateData())
	m.scrapeManager.ProvisionHotmailNewAccount()
	m.scrapeManager.Scrape()
	m.saveProfile(m.scrapeManager.ProfileData())
}

func (m *Master) saveProfile(profile map[string]string) {
	saver := files.NewManager()
	saver.SaveProfile(profile["email"], profile)

}

func (m *Master) generateData() map[string][]string {
	profile := m.profileManager.StandardProfile()
	profile.Generate()
	return profile.Values
}

func NewMaster() *Master {
	return &Master{
		profileManager: profile.NewManager(),
		scrapeManager:  scrape.NewManager(8081),
	}
}
