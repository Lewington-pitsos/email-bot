// Package datageneration contains structs that generate and store values.
// These values are later used to craete user profiles.
package vault

import (
	"email-bot/offline/helpers"
	"encoding/json"
	"go/build"
	"io/ioutil"
)

var dataPath = build.Default.GOPATH + "/src/email-bot/offline/datageneration/vault/data/"
var dataSuffix = ".json"

// +-------------------------------------------------------------------------------------+
// 									Vault STRUCT
// +-------------------------------------------------------------------------------------+

type Vault struct {
	bankFile    string
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (v *Vault) filePath(filename string) string {
	return dataPath + filename + dataSuffix
}

// Vault.allValues reads in the data file corresponding to the path/
// it decodes it from JSON
func (v *Vault) allValues(filename string) []string {
	filePath := v.filePath(filename)
	fileBytes, error := ioutil.ReadFile(filePath)

	helpers.Check(error)

	fileSlice := make([]string, 0, 10000)
	json.Unmarshal(fileBytes, &fileSlice)

	return fileSlice
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
		bankFile:    bankFile,
	}
}
