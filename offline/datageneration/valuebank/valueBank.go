// Package valuebank is just a bag for Vault objects
package valuebank

import (
	"email-bot/offline/datageneration/vault"
	"fmt"
)

// +-------------------------------------------------------------------------------------+
// 									Bank STRUCT
// +-------------------------------------------------------------------------------------+

type Bank struct {
	vaults map[string]vault.VaultInterface
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (b *Bank) AddSpecialVault(vaultName string) *Bank {
	switch vaultName {
	case "datevault":
		fmt.Println("lol")
		b.vaults[vaultName] = vault.NewDateVault()
	}

	return b
}

func (b *Bank) AddVault(vaultName string, bankFile string) *Bank {
	b.vaults[vaultName] = vault.NewVault(bankFile)
	return b
}

func (b *Bank) GiveValue(vaultName string) string {
	return b.vaults[vaultName].GiveValue()
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

func NewBank() *Bank {
	return &Bank{
		vaults: make(map[string]vault.VaultInterface),
	}
}
