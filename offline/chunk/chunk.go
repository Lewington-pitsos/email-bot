package chunk

// Chunk is a single instruction used by valueGenerator to generate part of a string.
// This string will ultimately make up part of a Profile.
// It has no methods, basically just a bag.
type Chunk struct {
	Mode     string
	Source   string
	Modified bool
	ModBank  string
}

func (c *Chunk) SetModification(modBank string) *Chunk {
	c.Modified = true
	c.ModBank = modBank

	return c
}

func NewChunk(mode string, source string) *Chunk {
	return &Chunk{
		Mode:     mode,
		Modified: false,
		Source:   source,
	}
}
