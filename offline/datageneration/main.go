package datageneration

import (
	"email-bot/offline/datageneration/valuebank"
	"email-bot/offline/datageneration/valuegenerator"
	"fmt"
)

func Test() {
	bank := valuebank.NewBank()
	bank.AddVault("username", "female-en")
	bank.AddVault("slang", "internet-slang")

	format := make([]*valuegenerator.ValueSpec, 0, 10) 
	spec := valuegenerator.NewValueSpec(false, "username", false, "slang")
	format = append(format, spec)
	spec2 := valuegenerator.NewValueSpec(false, "username", true, "slang")
	format = append(format, spec2)
	spec3 := valuegenerator.NewValueSpec(true, "@hotmail.com", false, "slang")
	format = append(format, spec3)

	vg := valuegenerator.NewValueGenerator("username", bank, format)
	fmt.Println(vg.Generate())
}