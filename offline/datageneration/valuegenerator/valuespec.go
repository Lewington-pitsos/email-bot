package valuegenerator

// ValueSpec is a single instruction used by valueGenerator to generate part of a string.
// This string will ultimately make up part of a Profile.
// It has no methods, basically justa a bag.
type ValueSpec struct {
	Literal  bool
	Output   string
	Modified bool
	ModBank string
}

// NewValueSpec is the only way to return a valueFormat outside datageneration
func NewValueSpec(literal bool, output string, modified bool, modbank string) *ValueSpec {
	return &ValueSpec{
		Literal:  literal,
		Output:   output,
		Modified: modified,
		ModBank: modbank,
	}
}
