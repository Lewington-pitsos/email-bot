package valuegenerator

import (
	"email-bot/offline/datageneration/valuegenerator/valuebank"
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

	b := valuebank.NewValueBank("word", bank, modBank)

	format := []*subValueSpec{
		newsubValueSpec(false, "word", true),
		newsubValueSpec(true, "@", false),
		newsubValueSpec(false, "word", false),
		newsubValueSpec(false, "word", true),
	}

	banks := make(map[string]*valuebank.ValueBank)
	banks[b.Name] = b

	nvg := NewvalueGenerator("username", banks, format)

	for i := 0; i < 20; i++ {
		fmt.Println(nvg.Generate())
	}

}
