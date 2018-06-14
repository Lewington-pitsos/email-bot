package datageneration

import (
	"email-bot/offline/datageneration/valuebank"
	"email-bot/offline/datageneration/valuegenerator"
	"email-bot/offline/datastructure"
	"fmt"
)

func Test() {
	bank := valuebank.NewBank().AddVault("username", "female-en").AddVault("slang", "internet-slang").AddSpecialVault("datevault")

	format := make([]*datastructure.ValueSpec, 0, 10)
	spec := &datastructure.ValueSpec{
		Literal:  false,
		Output:   "username",
		Modified: false,
		ModBank:  "slang",
	}
	format = append(format, spec)
	spec2 := &datastructure.ValueSpec{
		Literal:  false,
		Output:   "username",
		Modified: true,
		ModBank:  "slang",
	}
	format = append(format, spec2)
	spec3 := &datastructure.ValueSpec{
		Literal:  true,
		Output:   "@hotmail.com",
		Modified: true,
		ModBank:  "slang",
	}
	format = append(format, spec3)
	spec4 := &datastructure.ValueSpec{
		Literal:  false,
		Output:   "datevault",
		Modified: false,
		ModBank:  "slang",
	}
	format = append(format, spec4)

	vg := valuegenerator.NewValueGenerator("username", bank, format)
	fmt.Println(vg.Generate())

}
