package datastructures

// Signal is sent by scrapes to the relay.
type Signal struct {
	Name  string
	Value interface{}
}
