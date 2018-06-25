package profile

import "email-bot/offline/chunk"

// +-------------------------------------------------------------------------------------+
// 									Field STRUCT
// +-------------------------------------------------------------------------------------+

type Field struct {
	Name        string
	ValueNumber int
	Format      []*valuespec.ValueSpec
}

func(f *Field) AddChunck(chunk *Chunk) *Field {
	
}

func NewField(name string, valuesNeeded int) *Field {
	return &Field{
		Name:        name,
		ValueNumber: valuesNeeded,
		Format:      make([]*valuspec.valuesNeeded, 0, valuesNeeded)
	}
}
