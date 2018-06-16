package profile

import "email-bot/offline/datageneration/generator"

type Profile struct {
	Values    map[string][]string
	fields    []*field
	Generator *generator.ValueGenerator
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

func NewProfile(generator *generator.ValueGenerator) *Profile {
	values := make(map[string][]string)
	generator.SetValues(values)

	return &Profile{
		Values:    values,
		fields:    make([]*field, 0, 50),
		Generator: generator,
	}
}
