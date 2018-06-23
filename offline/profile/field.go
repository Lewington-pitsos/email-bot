package profile

import "email-bot/offline/valuespec"

// +-------------------------------------------------------------------------------------+
// 									field STRUCT
// +-------------------------------------------------------------------------------------+

type field struct {
	Name        string
	ValueNumber int
	Format      []*valuespec.ValueSpec
}

func NewField(name string, valuesNeeded int, format []*valuespec.ValueSpec) *field {
	return &field{
		Name:        name,
		ValueNumber: valuesNeeded,
		Format:      format,
	}
}
