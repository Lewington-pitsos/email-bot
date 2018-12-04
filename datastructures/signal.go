package datastructures

// Chunk is sent by scrapes to the relay.
type Chunk struct {
	Name  string
	Value interface{}
}
