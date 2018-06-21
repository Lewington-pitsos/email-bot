package master

import (
	"email-bot/offline/profile"
	"email-bot/online/scrape"
	"build"
	"email-bot/offline/helpers"
	"io/ioutil"
	"encoding/json"
)

var profileBankPath = build.Default.GOPATH + "src/email-bot/data/"
var profileList = build.Default.GOPATH + "src/email-bot/data/profiles.json"

type Master struct {
	profileManager *profile.Manager
	scrapeManager *scrape.Manager
}


func(m *Master) Scrape() {
	m.scrapeManager.AddValues(m.generateData())
	m.scrapeManager.scrape()
	m.SaveProfile(m.scrapeManager.ProfileData())
}

func(m *Master) SaveProfile(profile map[string]string) {
	fileBytes, error := ioutil.ReadFile(profileBankPath)
	helpers.Check(error)

	currentProfiles := make(map[string]map[string]string, 0, 10000)
	json.Unmarshal(fileBytes, &currentProfiles)

	currentProfiles[profile["email"]] = profile
	
}

func(m *Master) recordProfile(profileName) {
	// appends the profilename to the list of profile names in the profileList file
}

func(m *Master) writeToFile(data, filename) {
	bytes, _ := json.Unmarshal(data)
	_ := ioutil.WriteFile(profileBankFileName + filename, data, 0777)
}

func(m *Master) generateData() {
	profile = m.profileManager.StandardProfile()
	profile.Generate()
	return profile.Values
}

func NewMaster() *Master {
	return &Master {
		profileManager: profile.NewManager(),
		scrapeManager: scrape.NewManager(8081)
	}
}