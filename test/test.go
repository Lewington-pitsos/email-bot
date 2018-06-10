package test

import (
	"email-bot/offline/datageneration"
	"fmt"
)

// Test runs some tests so I can be sure go is going teh right thing
func Test() {
	bank := []string{
		"harry",
		"Mike",
		"Big",
		"Salmon",
		"winning",
		"fido",
		"kill",
		"master",
		"Sadie",
		"Ken",
		"Tiffany",
	}

	modBank := []string{
		"xx",
		"a",
		"oOo",
		"3",
		"lol",
		"s",
		"z",
		"88",
		"1234",
	}

	b := datageneration.NewValueBank(bank, modBank)

	for i := 0; i < 20; i++ {
		fmt.Println(b.GiveValue())
	}

	for i := 0; i < 20; i++ {
		fmt.Println(b.GiveModifiedValue())
	}
}
