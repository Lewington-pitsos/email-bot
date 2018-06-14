package datastructure

// ValueSpec is a single instruction used by valueGenerator to generate part of a string.
// This string will ultimately make up part of a Profile.
// It has no methods, basically justa a bag.
type ValueSpec struct {
	Literal  bool
	Output   string
	Modified bool
	ModBank  string
}
