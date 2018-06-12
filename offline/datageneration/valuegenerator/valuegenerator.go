package valuegenerator

import "email-bot/offline/datageneration/valuebank"

// +-------------------------------------------------------------------------------------+
// 									ValueGenerator STRUCT
// +-------------------------------------------------------------------------------------+

// ValueGenerator generates a string value according to the instructions laid out in Format.
// ValueBanks contain sub-string which eventually combine into the returned string.
type ValueGenerator struct {
	Name   string
	bank  valuebank.Bank
	format []*subValueSpec
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

// Generate iterates over all the subValueSpec's in format.
// For each of these it generates a string, according to that spec.
// All the substrings are concatonated, and the result is returned.
func (vg *ValueGenerator) Generate() string {
	output := ""
	for _, spec := range vg.format {
		output += vg.getSubValue(spec)
	}

	return output
}

// getSubValue checks if the passed in subValueSpec is a litral.
// If so, it simply return's the spec's output field.
// Otherwise, it retrives the relevent (from vg.banks) and gets that bank to generate an output.
func (vg *ValueGenerator) getSubValue(svs *subValueSpec) string {
	if svs.Literal {
		return svs.Output
	} else {
		bank := vg.bank.GiveVault(svs.Output)
		if svs.Modified {
			return bank.GiveModifiedValue()
		} else {
			return bank.GiveValue()
		}
	}
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

// NewValueGenerator returns a ValueGenerator struct
func NewValueGenerator(
	name string,
	bank valuebank.Bank,
	format []*subValueSpec,
) *ValueGenerator {
	return &ValueGenerator{
		Name:   name,
		bank:   bank,
		format: format,
	}
}