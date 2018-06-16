package offline

import (
	"email-bot/offline/datageneration/generator"
	"email-bot/offline/datageneration/valuebank"
	"email-bot/offline/datastructure"
	"email-bot/offline/profile"
	"fmt"
)

var userProfile *profile.Profile

func Test() {
	bank := setupBank()
	generator := generator.NewValueGenerator(bank)
	userProfile = profile.NewProfile(generator)

	usernameFormat := setupUsernameFormat()
	usernameField := profile.NewField("username", 3, usernameFormat)
	userProfile.AddField(usernameField)

	emailFormat := setupEmailFormat()
	emailField := profile.NewField("email", 10, emailFormat)
	userProfile.AddField(emailField)

	dateFormat := setupDateFormat()
	dateField := profile.NewField("birthdate", 1, dateFormat)
	userProfile.AddField(dateField)

	userProfile.Generate()

	fmt.Println(userProfile.Values)
}

func setupBank() *valuebank.Bank {
	bank := valuebank.NewBank()
	bank.AddVault("username", "female-en")
	bank.AddVault("slang", "internet-slang")
	bank.AddSpecialVault("datevault")

	return bank
}

func setupUsernameFormat() []*datastructure.ValueSpec {
	usernameFormat := make([]*datastructure.ValueSpec, 0, 10)
	spec1 := datastructure.NewValueSpec(false, "username")
	spec2 := datastructure.NewValueSpec(false, "username")

	usernameFormat = append(usernameFormat, spec1)
	usernameFormat = append(usernameFormat, spec2)

	return usernameFormat
}

func setupEmailFormat() []*datastructure.ValueSpec {
	emailFormat := make([]*datastructure.ValueSpec, 0, 10)
	spec2 := datastructure.NewValueSpec(false, "username").SetModification("slang").SetProgenitor("username")
	spec3 := datastructure.NewValueSpec(true, "@hotmail.com")

	emailFormat = append(emailFormat, spec2)
	emailFormat = append(emailFormat, spec3)

	return emailFormat
}

func setupDateFormat() []*datastructure.ValueSpec {
	dateFormat := make([]*datastructure.ValueSpec, 0, 10)
	spec1 := datastructure.NewValueSpec(false, "datevault")

	dateFormat = append(dateFormat, spec1)

	return dateFormat
}
