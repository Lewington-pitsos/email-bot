package datageneration

// SubValueSpec is a single instruction used by valueGenerator to generate part of a string.
// This string will ultimately make up part of a Profile.
// It has no methods, basically justa a bag.
type SubValueSpec struct {
	Literal  bool
	Output   string
	Modified bool
}

// NewSubValueSpec is the only way to return a valueFormat outside datageneration
func NewSubValueSpec(literal bool, output string, modified bool) *SubValueSpec {
	return &SubValueSpec{
		Literal:  literal,
		Output:   output,
		Modified: modified,
	}
}
