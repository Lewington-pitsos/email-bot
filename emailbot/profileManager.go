package emailbot

import (
	"email-bot/offline/profile"
)

type ProfileManager struct {
}

func (pm *ProfileManager) NewDataProfile() *profile.DataProfile {
	return profile.NewDataProfile()
}

func (pm *ProfileManager) NewExistingProfile(profileName string) *profile.ActiveProfile {
	return profile.NewActiveProfile(profileName)
}

func NewProfileManager() *ProfileManager {
	return &ProfileManager{}
}
