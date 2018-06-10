package datageneration

// valueGenerator generates a string value according to the instructions laid out in valueFormat.
// valueBanks contain sub-string which eventually combine into the returned string.
type valueGenerator struct {
	name        string
	banks       map[string]valueBank
	valueFormat []valueFormat
}
