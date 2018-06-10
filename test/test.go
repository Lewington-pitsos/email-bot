package test

import (
	"email-bot/offline/detailType"
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

	b := detailType.NewValueBank(bank, modBank)

	for i := 0; i < 20; i++ {
		fmt.Println(b.GiveValue())
	}

	for i := 0; i < 20; i++ {
		fmt.Println(b.GiveModifiedValue())
	}
}
