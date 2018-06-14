package offline

import (
	"email-bot/offline/datageneration/valuebank"
	"email-bot/offline/datageneration/valuegenerator"
	"email-bot/offline/datastructure"
	"email-bot/offline/profile"
	"fmt"
)

var userProfile *profile.Profile

func Test() {
	bank := setupBank()
	userProfile = profile.NewProfile()

	usernameFormat := setupUsernameFormat()
	usernameGenerator := valuegenerator.NewValueGenerator(bank, usernameFormat)
	usernameField := profile.NewField("username", 3, usernameGenerator)
	userProfile.AddField(usernameField)

	emailFormat := setupEmailFormat()
	emailGenerator := valuegenerator.NewValueGenerator(bank, emailFormat)
	emailField := profile.NewField("email", 100, emailGenerator)
	userProfile.AddField(emailField)

	dateFormat := setupDateFormat()
	dateGenerator := valuegenerator.NewValueGenerator(bank, dateFormat)
	dateField := profile.NewField("birthdate", 1, dateGenerator)
	userProfile.AddField(dateField)

	userProfile.Generate()

	fmt.Println(userProfile.Profile())
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
