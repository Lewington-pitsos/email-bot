// Package valuebank is just a bag for Vault objects
package valuebank

import (
	"email-bot/offline/vault"
	"go/build"
)

const bankFileSuffix = ".json"
var bankPath string = build.Default.GOPATH + "/src/email-bot/data/bankvalues/"

// +-------------------------------------------------------------------------------------+
// 									Bank STRUCT
// +-------------------------------------------------------------------------------------+

type Bank struct {
	vaults map[string]vault.VaultInterface
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (b *Bank) bankPath(bankName string) string {
	return bankPath + bankName + bankFileSuffix
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (b *Bank) AddVault(vaultName string, sourceName string) *Bank {
	switch vaultName {
	case "datevault":
		b.vaults[vaultName] = vault.NewDateVault(sourceName)
	case "passvault":
		b.vaults[vaultName] = vault.NewPassVault()
	}

	b.vaults[vaultName] = vault.NewVault(b.bankPath(sourceName))
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
