package profile

// import (
// 	"email-bot/logger"
// 	"email-bot/offline/generator"
// 	"email-bot/offline/valuespec"
// )

// type Designer struct {
// }

// func (m *Designer) StandardProfile() *Profile {
// 	logger.LoggerInterface.Println("Generating standard profile")
// 	generator := generator.NewValueGenerator()
// 	profile := NewProfile(generator)

// 	firstNameField := NewField("firstname", 1, []*valuespec.ValueSpec{
// 		valuespec.NewValueSpec("bank", "name"),
// 	})

// 	lastNameField := NewField("lastname", 1, []*valuespec.ValueSpec{
// 		valuespec.NewValueSpec("bank", "name"),
// 	})

// 	fullNameField := NewField("fullname", 1, []*valuespec.ValueSpec{
// 		valuespec.NewValueSpec("derived", "firstname"),
// 		valuespec.NewValueSpec("derived", "lastname"),
// 	})

// 	emailField := NewField("email", 100, []*valuespec.ValueSpec{
// 		valuespec.NewValueSpec("derived", "fullname").SetModification("slang"),
// 		valuespec.NewValueSpec("literal", "@hotmail.com"),
// 	})

// 	dayField := NewField("day", 1, []*valuespec.ValueSpec{
// 		valuespec.NewValueSpec("bank", "dayvault"),
// 	})

// 	monthField := NewField("month", 1, []*valuespec.ValueSpec{
// 		valuespec.NewValueSpec("bank", "monthvault"),
// 	})

// 	yearField := NewField("year", 1, []*valuespec.ValueSpec{
// 		valuespec.NewValueSpec("bank", "yearvault"),
// 	})

// 	passField := NewField("password", 20, []*valuespec.ValueSpec{
// 		valuespec.NewValueSpec("bank", "passvault"),
// 	})

// 	profile.AddFields(firstNameField, lastNameField, fullNameField, emailField, dayField, monthField, yearField, passField)

// 	return profile
// }

// func NewDesigner() *Designer {
// 	return &Designer{}
// }
