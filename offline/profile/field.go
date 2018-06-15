package profile

import "email-bot/offline/datastructure"

// +-------------------------------------------------------------------------------------+
// 									field STRUCT
// +-------------------------------------------------------------------------------------+

type field struct {
	Name        string
	ValueNumber int
	Format      []*datastructure.ValueSpec
}

func NewField(name string, valuesNeeded int, format []*datastructure.ValueSpec) *field {
	return &field{
		Name:        name,
		ValueNumber: valuesNeeded,
		Format:      format,
	}
}
