package profile

import (
	"email-bot/datastructures"
	"email-bot/logger"
	"email-bot/offline/generator"
)

type DataProfile struct {
	Values    map[string]datastructures.Detail
	fields    []*Field
	Generator *generator.ValueGenerator
}

func (p *DataProfile) AddField(name string, ValueNumber int) *Field {
	logger.LoggerInterface.Println("Adding field:", field.Name)
	field := NewField(name, ValueNumber)
	p.fields = append(p.fields, field)

	return field
}


func (p *DataProfile) Generate() *DataProfile {
	logger.LoggerInterface.Println("Generating DataProfile data")
	for _, field := range p.fields {
		p.Values[field.Name] = p.Generator.Generate(field.Format, field.ValueNumber)
	}

	return p
}

func NewDataProfile() *DataProfile {
	generator := generator.NewValueGenerator()
	logger.LoggerInterface.Println("Creating DataProfile")
	values := make(map[string]datastructures.Detail)
	generator.SetValues(values)

	return &DataProfile{
		Values:    values,
		fields:    make([]*field, 0, 50),
		Generator: generator,
	}
}
