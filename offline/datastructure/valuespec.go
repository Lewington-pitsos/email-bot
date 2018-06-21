package datastructure

// ValueSpec is a single instruction used by valueGenerator to generate part of a string.
// This string will ultimately make up part of a Profile.
// It has no methods, basically justa a bag.
type ValueSpec struct {
	Mode 	 string
	Output   string
	Modified bool
	ModBank  string
}

func (v *ValueSpec) SetModification(modBank string) *ValueSpec {
	v.Modified = true
	v.ModBank = modBank

	return v
}

func NewValueSpec(mode string, output string) *ValueSpec {
	return &ValueSpec{
		Mode: mode,
		Modified: false,
		Output:   output,
	}
}
