package valuegenerator

import (
	"email-bot/offline/datageneration/valuebank"
	"email-bot/offline/datastructure"
	"email-bot/offline/helpers"
)

// +-------------------------------------------------------------------------------------+
// 									ValueGenerator STRUCT
// +-------------------------------------------------------------------------------------+

// ValueGenerator generates a string value according to the instructions laid out in Format.
// ValueBanks contain sub-string which eventually combine into the returned string.
type ValueGenerator struct {
	Name   string
	bank   *valuebank.Bank
	format []*datastructure.ValueSpec
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

// Inserss the insert into the string at index and returns the result
// Doesn't add back the remaining half of str if index is too close to the end of str
func (vg *ValueGenerator) modifiedString(str string, insert string, index int) string {
	if index <= len(str)-1 {
		return str[:index] + insert + str[index+1:]
	} else {
		return str[:index] + insert
	}
}

func (vg *ValueGenerator) modifiedValue(value string, modification string) string {
	modIndex := helpers.GetRandomIndex(value)
	return vg.modifiedString(value, modification, modIndex)
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

// Generate iterates over all the ValueSpec's in format.
// For each of these it generates a string, according to that spec.
// All the substrings are concatonated, and the result is returned.
func (vg *ValueGenerator) Generate() string {
	output := ""
	for _, spec := range vg.format {
		output += vg.getSubValue(spec)
	}

	return output
}

// getSubValue checks if the passed in ValueSpec is a litral.
// If so, it simply return's the spec's output field.
// Otherwise, it retrives the relevent (from vg.banks) and gets that bank to generate an output.
func (vg *ValueGenerator) getSubValue(svs *datastructure.ValueSpec) string {
	if svs.Literal {
		return svs.Output
	} else {
		value := vg.bank.GiveValue(svs.Output)
		if svs.Modified {
			modification := vg.bank.GiveValue(svs.ModBank)
			return vg.modifiedValue(value, modification)
		} else {
			return value
		}
	}
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

// NewValueGenerator returns a ValueGenerator struct
func NewValueGenerator(
	name string,
	bank *valuebank.Bank,
	format []*datastructure.ValueSpec,
) *ValueGenerator {
	return &ValueGenerator{
		Name:   name,
		bank:   bank,
		format: format,
	}
}
