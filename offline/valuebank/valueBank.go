// Package valuebank is just a bag for Vault objects
package valuebank

import (
	"email-bot/offline/vault"
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
		b.vaults[vaultName] = vault.NewDateVault("02/01/2006")
	case "yearvault":
		b.vaults[vaultName] = vault.NewDateVault("2006")
	case "monthvault":
		b.vaults[vaultName] = vault.NewDateVault("January")
	case "dayvault":
		b.vaults[vaultName] = vault.NewDateVault("02")
	case "passvault":
		b.vaults[vaultName] = vault.NewPassVault()
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
