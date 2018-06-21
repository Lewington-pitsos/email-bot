// Package contains structs that generate and store values.
// These values are later used to craete user profiles.
package vault

import (
	"email-bot/offline/files"
	"email-bot/offline/helpers"
)

// +-------------------------------------------------------------------------------------+
// 									Vault STRUCT
// +-------------------------------------------------------------------------------------+

type Vault struct {
	bankFile string
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

// Vault.allValues reads in the data file corresponding to the path/
// it decodes it from JSON
func (v *Vault) allValues(filename string) []string {
	loader := files.NewManager()

	return loader.BankeData(v.bankFile)
}

func (v *Vault) randomValue(valueFileName string) string {
	values := v.allValues(valueFileName)
	return helpers.GetRandom(values)
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (v *Vault) GiveValue() string {
	return v.randomValue(v.bankFile)
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

// NewVault returns a new Vault struct.
// It's the only way to get hold of a Vault struct outside valueGenerator.
func NewVault(bankFile string) *Vault {
	return &Vault{
		bankFile: bankFile,
	}
}
