package main

import (
	"email-bot/offline/datageneration/generator"
	"email-bot/offline/datageneration/valuebank"
	"email-bot/offline/profile"
	"fmt"
)

func main() {
	bank := valuebank.SetupBank()
	generator := generator.NewValueGenerator(bank)
	manager := profile.NewManager()
	userProfile := manager.StandardProfile(generator)

	userProfile.Generate()

	values := userProfile.Values
	fmt.Println(values)
}
