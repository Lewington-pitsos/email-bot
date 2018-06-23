package profile

import (
	"email-bot/logger"
	"email-bot/offline/datastructure"
	"email-bot/offline/generator"
)

type Manager struct {
}

func (m *Manager) StandardProfile() *Profile {
	logger.LoggerInterface.Println("Generating standard profile")
	generator := generator.NewValueGenerator()
	profile := NewProfile(generator)

	firstNameField := NewField("firstname", 1, []*datastructure.ValueSpec{
		datastructure.NewValueSpec("bank", "name"),
	})

	lastNameField := NewField("lastname", 1, []*datastructure.ValueSpec{
		datastructure.NewValueSpec("bank", "name"),
	})

	fullNameField := NewField("fullname", 1, []*datastructure.ValueSpec{
		datastructure.NewValueSpec("derived", "firstname"),
		datastructure.NewValueSpec("derived", "lastname"),
	})

	emailField := NewField("email", 100, []*datastructure.ValueSpec{
		datastructure.NewValueSpec("derived", "fullname").SetModification("slang"),
		datastructure.NewValueSpec("literal", "@hotmail.com"),
	})

	dayField := NewField("day", 10, []*datastructure.ValueSpec{
		datastructure.NewValueSpec("bank", "dayvault"),
	})

	monthField := NewField("month", 10, []*datastructure.ValueSpec{
		datastructure.NewValueSpec("bank", "monthvault"),
	})

	yearField := NewField("year", 10, []*datastructure.ValueSpec{
		datastructure.NewValueSpec("bank", "yearvault"),
	})

	passField := NewField("password", 20, []*datastructure.ValueSpec{
		datastructure.NewValueSpec("bank", "passvault"),
	})

	profile.AddFields(firstNameField, lastNameField, fullNameField, emailField, dayField, monthField, yearField, passField)

	return profile
}

func NewManager() *Manager {
	return &Manager{}
}
