package chunk

// Chunk is a single instruction used by valueGenerator to generate part of a string.
// This string will ultimately make up part of a Profile.
// It has no methods, basically justa a bag.
type Chunk struct {
	Mode     string
	Output   string
	Modified bool
	ModBank  string
}

func (c *Chunk) SetModification(modBank string) *Chunk {
	c.Modified = true
	c.ModBank = modBank

	return c
}

func NewChunk(mode string, output string) *Chunk {
	return &Chunk{
		Mode:     mode,
		Modified: false,
		Output:   output,
	}
}
