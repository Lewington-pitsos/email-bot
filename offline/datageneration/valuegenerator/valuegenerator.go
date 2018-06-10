package valuegenerator

import "email-bot/offline/datageneration/valuegenerator/valuebank"

// valueGenerator generates a string value according to the instructions laid out in Format.
// ValueBanks contain sub-string which eventually combine into the returned string.
type valueGenerator struct {
	Name   string
	banks  map[string]*valuebank.ValueBank
	format []*subValueSpec
}

// NewValueGenerator returns a valuegGenerator struct
func NewvalueGenerator(
	name string,
	banks map[string]*valuebank.ValueBank,
	format []*subValueSpec,
) *valueGenerator {
	return &valueGenerator{
		Name:   name,
		banks:  banks,
		format: format,
	}
}

// Generate iterates over all the subValueSpec's in format.
// For each of these it generates a string, according to that spec.
// All the substrings are concatonated, and the result is returned.
func (vg *valueGenerator) Generate() string {
	output := ""
	for _, spec := range vg.format {
		output += vg.getSubValue(spec)
	}

	return output
}

// getSubValue checks if the passed in subValueSpec is a litral.
// If so, it simply return's the spec's output field.
// Otherwise, it retrives the relevent (from vg.banks) and gets that bank to generate an output.
func (vg *valueGenerator) getSubValue(svs *subValueSpec) string {
	if svs.Literal {
		return svs.Output
	} else {
		bank := vg.banks[svs.Output]
		if svs.Modified {
			return bank.GiveModifiedValue()
		} else {
			return bank.GiveValue()
		}
	}
}
