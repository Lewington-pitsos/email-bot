package profile

import (
	"email-bot/database"
	"email-bot/datastructures"
)

type ActiveProfile struct {
	Values map[string]datastructures.Detail
	Name   string
}

func (a *ActiveProfile) saveDetails(profileData map[string]string) {
	for key, value := range profileData {
		a.Values[key] = datastructures.NewDetaiMono(value)
	}
}

func (a *ActiveProfile) Populate() {
	loader := database.NewLibrarian()
	data := loader.FindProfiles(1)[0]
	loader.Close()
	a.saveDetails(data)
}

func NewActiveProfile(name string) *ActiveProfile {
	return &ActiveProfile{
		Values: make(map[string]datastructures.Detail),
		Name:   name,
	}
}
