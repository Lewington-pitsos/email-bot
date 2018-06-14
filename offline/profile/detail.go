package profile

import (
	"email-bot/offline/datageneration/valuegenerator"
)

// +-------------------------------------------------------------------------------------+
// 									field STRUCT
// +-------------------------------------------------------------------------------------+

type field struct {
	Name      string
	Values    []string
	generator *valuegenerator.ValueGenerator
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (d *field) Generate() {
	for i := 0; i < cap(d.Values); i++ {
		d.Values[i] = d.generator.Generate()
	}
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

func Newfield(name string, valuesNeeded int, generator *valuegenerator.ValueGenerator) *field {
	return &field{
		Name:      name,
		Values:    make([]string, valuesNeeded),
		generator: generator,
	}
}
