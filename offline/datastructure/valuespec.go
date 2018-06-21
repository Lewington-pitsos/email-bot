package datastructure

// ValueSpec is a single instruction used by valueGenerator to generate part of a string.
// This string will ultimately make up part of a Profile.
// It has no methods, basically justa a bag.
type ValueSpec struct {
	Literal  bool
	Output   string
	Modified bool
	ModBank  string
	Derived  bool
}

func (v *ValueSpec) SetModification(modBank string) *ValueSpec {
	v.Modified = true
	v.ModBank = modBank

	return v
}

func (v *ValueSpec) SetProgenitor(progenitor string) *ValueSpec {
	v.Derived = true
	v.Output = progenitor

	return v
}

func NewValueSpec(literal bool, output string) *ValueSpec {
	return &ValueSpec{
		Literal:  literal,
		Output:   output,
		Modified: false,
	}
}
