package datageneration

import (
	"email-bot/offline/datageneration/valuebank"
)

func Test() {
	bank := valuebank.NewBank()
	bank.AddVault("username", "female-en", "internet-slang")

	v1 := bank.GiveVault("username")

	for {
		v1.GiveValue()
	}
}