package datageneration

import (
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

	b := NewvalueBank("word", bank, modBank)

	format := []*SubValueSpec{
		NewSubValueSpec(false, "word", true),
		NewSubValueSpec(true, "@", false),
		NewSubValueSpec(false, "word", false),
		NewSubValueSpec(false, "word", true),
	}

	banks := make(map[string]*valueBank)
	banks[b.Name] = b

	nvg := NewvalueGenerator("username", banks, format)

	for i := 0; i < 20; i++ {
		fmt.Println(nvg.Generate())
	}

}
