package profile

import "email-bot/offline/chunk"

// +-------------------------------------------------------------------------------------+
// 									Field STRUCT
// +-------------------------------------------------------------------------------------+

type Field struct {
	Name        string
	ValueNumber int
	Format      []*chunk.Chunk
}

func NewField(name string, valuesNeeded int) *Field {
	return &Field{
		Name:        name,
		ValueNumber: valuesNeeded,
		Format:      make([]*chunk.Chunk, 0, valuesNeeded),
	}
}
