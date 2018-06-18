package main

import (
	"email-bot/offline/datageneration/generator"
	"email-bot/offline/datageneration/valuebank"
	"email-bot/offline/profile"
	"email-bot/online/scrape"
	"fmt"
	"time"
)

func main() {
	bank := valuebank.SetupBank()
	generator := generator.NewValueGenerator(bank)
	manager := profile.NewManager()
	userProfile := manager.StandardProfile(generator)

	userProfile.Generate()

	values := userProfile.Values
	fmt.Println(values)

	m := scrape.NewManager(8083)
	m.AddValues(valies)
	m.ProvisionHotmailNewAccount()
	m.Scrape()

	time.Sleep(time.Millisecond * 80000)
}
