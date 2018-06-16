package profile

import (
	"email-bot/offline/datageneration/generator"
	"email-bot/offline/datastructure"
)

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) StandardProfile(generator *generator.ValueGenerator) *Profile {
	profile := NewProfile(generator)

	usernameField := NewField("username", 3, m.usernameFormat())
	emailField := NewField("email", 100, m.emailFormat())
	dateField := NewField("birthdate", 1, m.dateFormat())

	profile.AddFields(usernameField, emailField, dateField)

	return profile
}

func (m *Manager) usernameFormat() []*datastructure.ValueSpec {
	usernameFormat := make([]*datastructure.ValueSpec, 0, 10)
	spec1 := datastructure.NewValueSpec(false, "username")
	spec2 := datastructure.NewValueSpec(false, "username")

	usernameFormat = append(usernameFormat, spec1)
	usernameFormat = append(usernameFormat, spec2)

	return usernameFormat
}

func (m *Manager) emailFormat() []*datastructure.ValueSpec {
	emailFormat := make([]*datastructure.ValueSpec, 0, 10)
	spec2 := datastructure.NewValueSpec(false, "username").SetModification("slang").SetProgenitor("username")
	spec3 := datastructure.NewValueSpec(true, "@hotmail.com")

	emailFormat = append(emailFormat, spec2)
	emailFormat = append(emailFormat, spec3)

	return emailFormat
}

func (m *Manager) dateFormat() []*datastructure.ValueSpec {
	dateFormat := make([]*datastructure.ValueSpec, 0, 10)
	spec1 := datastructure.NewValueSpec(false, "datevault")

	dateFormat = append(dateFormat, spec1)

	return dateFormat
}
