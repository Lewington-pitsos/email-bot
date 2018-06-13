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

func (b *Bank) AddVault(vaultName string, bankFile string) {
	b.vaults[vaultName] = vault.NewVault(bankFile)
}

func (b *Bank) GiveValue(vaultName string) string {
	return b.vaults[vaultName].GiveValue()
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

func NewBank() *Bank {
	return &Bank {
		vaults: make(map[string]*vault.Vault),
	}
}
