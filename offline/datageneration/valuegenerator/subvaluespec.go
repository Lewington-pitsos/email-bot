package valuegenerator

// subValueSpec is a single instruction used by valueGenerator to generate part of a string.
// This string will ultimately make up part of a Profile.
// It has no methods, basically justa a bag.
type subValueSpec struct {
	Literal  bool
	Output   string
	Modified bool
}

// NewsubValueSpec is the only way to return a valueFormat outside datageneration
func newsubValueSpec(literal bool, output string, modified bool) *subValueSpec {
	return &subValueSpec{
		Literal:  literal,
		Output:   output,
		Modified: modified,
	}
}
