package profile

import (
	"email-bot/database"
	"email-bot/datastructures"
)

type ActiveProfile struct {
	values map[string]datastructures.Detail
	Name   string
}

func (a *ActiveProfile) saveDetails(profileData map[string]string) {
	for key, value := range profileData {
		a.values[key] = datastructures.NewDetaiMono(value)
	}
}

func(a *ActiveProfile) Values() map[string]datastructures.Detail {
	return a.values
}

func (a *ActiveProfile) Generate() {
	loader := database.NewLibrarian()
	data := loader.FindProfiles(1)[0]
	loader.Close()
	a.saveDetails(data)
}

func NewActiveProfile(name string) *ActiveProfile {
	return &ActiveProfile{
		values: make(map[string]datastructures.Detail),
		Name:   name,
	}
}
