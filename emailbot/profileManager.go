package emailbot

import (
	"email-bot/offline/profile"
)

type ProfileManager struct {
}

func (pm *ProfileManager) NewDataProfile() *profile.dataProfile {
	return profile.NewProfile()
}

func (pm *ProfileManager) NewExistingProfile(profileName string) *profile.activeProfile {
	return profile.NewActiveProfile(profileName)
}

func NewProfileManager(dataProfile bool) *ProfileManager {
	return &ProfileManager {
	}
}

