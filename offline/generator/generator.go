package generator

import (
	"email-bot/datastructures"
	"email-bot/offline/chunk"
	"email-bot/offline/helpers"
	"email-bot/offline/valuebank"
)

// +-------------------------------------------------------------------------------------+
// 									ValueGenerator STRUCT
// +-------------------------------------------------------------------------------------+

// ValueGenerator generates a string value according to the instructions laid out in Format.
// ValueBanks contain sub-string which eventually combine into the returned string.
type ValueGenerator struct {
	bank   *valuebank.Bank
	values map[string]datastructures.Detail
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

// getSubValue checks if the passed in Chunk is a litral.
// If so, it simply return's the spec's output field.
// Otherwise, it retrives the relevent (from vg.banks) and gets that bank to generate an output.
func (vg *ValueGenerator) getSubValue(svs *chunk.Chunk) string {
	unmodifiedValue := vg.unmodifiedValue(svs)

	if svs.Modified {
		modification := vg.bank.GiveValue(svs.ModBank)
		return vg.modifiedValue(unmodifiedValue, modification)
	}

	return unmodifiedValue
}

func (vg *ValueGenerator) unmodifiedValue(svs *chunk.Chunk) string {
	switch svs.Mode {
	case "literal":
		return svs.Source
	case "derived":
		return vg.values[svs.Source].RandomValue()
	case "bank":
		return vg.bank.GiveValue(svs.Source)
	}

	panic("invalid mode")
	return "invalid mode"
}

func (vg *ValueGenerator) generateValue(format []*chunk.Chunk) string {
	value := ""
	for _, spec := range format {
		value += vg.getSubValue(spec)
	}

	return value
}

func (vg *ValueGenerator) detailMany(format []*chunk.Chunk, valueNumber int) datastructures.Detail {
	values := make([]string, valueNumber)

	for i := 0; i < valueNumber; i++ {
		values[i] = vg.generateValue(format)
	}

	return datastructures.NewDetaiMany(values)
}

func (vg *ValueGenerator) detailMono(format []*chunk.Chunk) datastructures.Detail {
	return datastructures.NewDetaiMono(vg.generateValue(format))
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (vg *ValueGenerator) SetValues(values map[string]datastructures.Detail) *ValueGenerator {
	vg.values = values
	return vg
}

// Generate iterates over all the Chunk's in format.
// For each of these it generates a string, according to that spec.
// All the substrings are concatonated, and the result is returned.
func (vg *ValueGenerator) Generate(format []*chunk.Chunk, valueNumber int) datastructures.Detail {
	if valueNumber == 1 {
		return vg.detailMono(format)
	}

	return vg.detailMany(format, valueNumber)
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

// NewValueGenerator returns a ValueGenerator struct
func NewValueGenerator() *ValueGenerator {
	return &ValueGenerator{
		bank: valuebank.SetupBank(),
	}
}
