package profile

import (
	"email-bot/offline/generator"	
	"email-bot/logger"

)
type Profile struct {
	Values    map[string][]string
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

func (p *Profile) Generate() *Profile {
	logger.LoggerInterface.Println("Generating Profile data")
	for _, field := range p.fields {
		p.Values[field.Name] = make([]string, field.ValueNumber)
		for i := 0; i < field.ValueNumber; i++ {
			p.Values[field.Name][i] = p.Generator.Generate(field.Format)
		}
	}

	return p
}

func NewProfile(generator *generator.ValueGenerator) *Profile {
	logger.LoggerInterface.Println("Creating Profile")
	values := make(map[string][]string)
	generator.SetValues(values)

	return &Profile{
		Values:    values,
		fields:    make([]*field, 0, 50),
		Generator: generator,
	}
}
