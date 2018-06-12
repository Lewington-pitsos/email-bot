// Package valuebank is just a bag for Vault objects
package valuebank

import "email-bot/offline/datageneration/vault"

// +-------------------------------------------------------------------------------------+
// 									Bank STRUCT
// +-------------------------------------------------------------------------------------+


type Bank struct {
	vaults map[string]*vault.Vault
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (b *Bank) AddVault(vaultName string, bankFile string, modFile string) {
	b.vaults[vaultName] = vault.NewVault(bankFile, modFile)
}

func (b *Bank) GiveVault(vaultName string) *vault.Vault {
	return b.vaults[vaultName]
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

func NewBank() *Bank {
	return &Bank {
		vaults: make(map[string]*vault.Vault),
	}
}
