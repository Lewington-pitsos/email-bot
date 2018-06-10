// Package datageneration contains structs that generate and store values.
// These values are later used to craete user profiles.
package valuebank

import "email-bot/offline/helpers"

// +-------------------------------------------------------------------------------------+
// 							DEFINING ValueBank STRUCT
// +-------------------------------------------------------------------------------------+

type ValueBank struct {
	Name    string
	bank    []string
	modBank []string
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

// Returns GiveValue() except a string from v.modBank is inserted in palace of one of the
// characters
func (v *ValueBank) GiveModifiedValue() string {
	selectedString := v.GiveValue()
	modIndex := helpers.GetRandomIndex(selectedString)
	modString := v.giveModValue()

	return v.makeModifiedString(selectedString, modString, modIndex)
}

// Insers the insert into the string at index and returns the result
// Doesn't add back the remaining half of str if index is too close to the end of str
func (v *ValueBank) makeModifiedString(str string, insert string, index int) string {
	if index <= len(str)-1 {
		return str[:index] + insert + str[index+1:]
	} else {
		return str[:index] + insert
	}
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (v *ValueBank) GiveValue() string {
	return helpers.GetRandom(v.bank)
}

func (v *ValueBank) giveModValue() string {
	return helpers.GetRandom(v.modBank)
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

// NewValueBank returns a new ValueBank struct.
// It's the only way to get hold of a ValueBank struct outside valueGenerator.
func NewValueBank(name string, bankValues []string, modBankValues []string) *ValueBank {
	return &ValueBank{
		Name:    name,
		bank:    bankValues,
		modBank: modBankValues,
	}
}
