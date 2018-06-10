package datageneration

// valueFormat is a single instruction used by valueGenerator to generate part of a string.
// This string will ultimately make up part of a Profile
type valueFormat struct {
	Literal bool
	Output  string
}
