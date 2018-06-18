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

	usernameField := NewField("username", 10, m.usernameFormat())
	emailField := NewField("email", 100, m.emailFormat())
	dayField := NewField("day", 10, m.dayFormat())
	monthField := NewField("month", 10, m.monthFormat())
	yearField := NewField("year", 10, m.yearFormat())
	passField := NewField("password", 20, m.passFormat())

	profile.AddFields(usernameField, emailField, dayField, monthField, yearField, passField)

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

func (m *Manager) yearFormat() []*datastructure.ValueSpec {
	yearFormat := make([]*datastructure.ValueSpec, 0, 10)
	spec1 := datastructure.NewValueSpec(false, "yearvault")

	yearFormat = append(yearFormat, spec1)

	return yearFormat
}

func (m *Manager) dayFormat() []*datastructure.ValueSpec {
	dayFormat := make([]*datastructure.ValueSpec, 0, 10)
	spec1 := datastructure.NewValueSpec(false, "dayvault")

	dayFormat = append(dayFormat, spec1)

	return dayFormat
}

func (m *Manager) monthFormat() []*datastructure.ValueSpec {
	monthFormat := make([]*datastructure.ValueSpec, 0, 10)
	spec1 := datastructure.NewValueSpec(false, "monthvault")

	monthFormat = append(monthFormat, spec1)

	return monthFormat
}

func (m *Manager) passFormat() []*datastructure.ValueSpec {
	passFormat := make([]*datastructure.ValueSpec, 0, 10)
	spec1 := datastructure.NewValueSpec(false, "passvault")

	passFormat = append(passFormat, spec1)

	return passFormat
}
