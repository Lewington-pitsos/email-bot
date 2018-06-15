package profile

import "email-bot/offline/datageneration/valuebank"

type Profile struct {
	Values    map[string][]string
	fields    []*field
	Generator *ValueGenerator
}

func (p *Profile) AddField(d *field) *Profile {
	p.fields = append(p.fields, d)

	return p
}

func (p *Profile) Generate() *Profile {
	for _, field := range p.fields {
		p.Values[field.Name] = make([]string, field.ValueNumber)
		for i := 0; i < field.ValueNumber; i++ {
			p.Values[field.Name][i] = p.Generator.Generate(field.Format)
		}
	}

	return p
}

func NewProfile(bank *valuebank.Bank) *Profile {
	values := make(map[string][]string)

	return &Profile{
		Values:    values,
		fields:    make([]*field, 0, 50),
		Generator: NewValueGenerator(bank, values),
	}
}
