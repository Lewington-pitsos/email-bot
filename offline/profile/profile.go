package profile

import "email-bot/offline/datageneration/valuebank"

type Profile struct {
	values    *map[string][]string
	fields    []*field
	Generator *ValueGenerator
}

func (p *Profile) AddField(d *field) *Profile {
	p.fields = append(p.fields, d)

	return p
}

func NewProfile(name string, bank *valuebank.Bank) *Profile {
	values := make(map[string][]string)

	return &Profile{
		values:    &values,
		fields:    make([]*field, 0, 50),
		Generator: NewValueGenerator(bank, &values),
	}
}
