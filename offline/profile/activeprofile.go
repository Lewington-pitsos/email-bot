package profile

import (
	"email-bot/database"
	"email-bot/datastructures"
)

type ActiveProfile struct {
	profiles     []map[string]datastructures.Detail
	profileIndex int
}

func (ap *ActiveProfile) Values() map[string]datastructures.Detail {
	return ap.profiles[ap.profileIndex]
}

func (ap *ActiveProfile) Generate() ProfileInterface {
	ap.profileIndex++

	return ap
}

func (a *ActiveProfile) Profiles() []map[string]datastructures.Detail {
	return a.profiles
}

func (a *ActiveProfile) Saveable() bool {
	return false
}

func (ap *ActiveProfile) Populate() *ActiveProfile {
	loader := database.NewLibrarian()
	data := loader.FindProfiles(20)
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

func NewActiveProfile() *ActiveProfile {
	return &ActiveProfile{
		profiles:     make([]map[string]datastructures.Detail, 0, 100),
		profileIndex: 0,
	}
}
