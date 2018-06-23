package profile

import (
	"email-bot/datastructures"
	"email-bot/logger"
	"email-bot/offline/generator"
)

type Profile struct {
	Values    map[string]datastructures.Detail
	fields    []*field
	Generator *generator.ValueGenerator
}

func (p *Profile) AddFields(fields ...*field) *Profile {
	for _, field := range fields {
		logger.LoggerInterface.Println("Adding field:", field.Name)
		p.fields = append(p.fields, field)
	}

	return p
}
func (p *Profile) detail(fields ...*field) *Profile {
	for _, field := range fields {
		logger.LoggerInterface.Println("Adding field:", field.Name)
		p.fields = append(p.fields, field)
	}

	return p
}

func (p *Profile) Generate() *Profile {
	logger.LoggerInterface.Println("Generating Profile data")
	for _, field := range p.fields {
		p.Values[field.Name] = p.Generator.Generate(field.Format, field.ValueNumber)
	}

	return p
}

func NewProfile(generator *generator.ValueGenerator) *Profile {
	logger.LoggerInterface.Println("Creating Profile")
	values := make(map[string]datastructures.Detail)
	generator.SetValues(values)

	return &Profile{
		Values:    values,
		fields:    make([]*field, 0, 50),
		Generator: generator,
	}
}
