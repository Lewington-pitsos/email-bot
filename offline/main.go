package offline

import (
	"email-bot/offline/generator"
	"email-bot/offline/valuebank"
	"email-bot/offline/profile"
	"fmt"
)

var userProfile *profile.Profile

func Test() {
	bank := valuebank.SetupBank()
	generator := generator.NewValueGenerator(bank)
	manager := profile.NewManager()
	userProfile := manager.StandardProfile(generator)

	userProfile.Generate()

	fmt.Println(userProfile.Values)
}
