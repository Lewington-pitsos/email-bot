package profile

import (
	"email-bot/database"
	"email-bot/datastructures"
)

type ActiveProfile struct {
	profiles     []map[string]datastructures.Detail
	profileIndex int
	Name         string
}

func (ap *ActiveProfile) Values() map[string]datastructures.Detail {
	return ap.profiles[ap.profileIndex]
}

func (ap *ActiveProfile) Generate() *ActiveProfile {
	ap.profileIndex++

	return ap
}

func (a *ActiveProfile) Profiles() []map[string]datastructures.Detail {
	return a.profiles
}

func (ap *ActiveProfile) Populate() *ActiveProfile {
	loader := database.NewLibrarian()
	data := loader.FindProfiles(1)
	loader.Close()
	for _, profile := range data {
		permanentProfile := make(map[string]datastructures.Detail)
		for key, value := range profile {
			permanentProfile[key] = datastructures.NewDetaiMono(value)
		}
		ap.profiles = append(ap.profiles, permanentProfile)
	}

	return ap
}

func NewActiveProfile(name string) *ActiveProfile {
	return &ActiveProfile{
		profiles:     make([]map[string]datastructures.Detail, 0, 100),
		profileIndex: 0,
		Name:         name,
	}
}
