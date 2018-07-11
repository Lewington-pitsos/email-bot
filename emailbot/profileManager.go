package emailbot

import (
	"email-bot/offline/profile"
)

type ProfileManager struct {
}

func (pm *ProfileManager) NewDataProfile() *profile.DataProfile {
	return profile.NewDataProfile()
}

func (pm *ProfileManager) NewExistingProfile() *profile.ActiveProfile {
	return profile.NewActiveProfile()
}

func NewProfileManager() *ProfileManager {
	return &ProfileManager{}
}
