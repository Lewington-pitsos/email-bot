package profile

import (
	"email-bot/datastructures"
	"email-bot/logger"
	"email-bot/offline/chunk"
	"email-bot/offline/generator"
)

type DataProfile struct {
	Values    map[string]datastructures.Detail
	fields    []*Field
	Generator *generator.ValueGenerator
}

func (p *DataProfile) AddField(name string, ValueNumber int) *DataProfile {
	field := NewField(name, ValueNumber)
	p.fields = append(p.fields, field)

	return p
}

func (p *DataProfile) WithChunk(mode string, source string) *DataProfile {
	field := p.fields[len(p.fields)-1]
	field.Format = append(field.Format, chunk.NewChunk(mode, source))

	return p
}

func (p *DataProfile) WithModifiedChunk(mode string, source string, modBank string) *DataProfile {
	modChunk := chunk.NewChunk(mode, source).SetModification(modBank)
	field := p.fields[len(p.fields)-1]
	field.Format = append(field.Format, modChunk)

	return p
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
		fields:    make([]*Field, 0, 50),
		Generator: generator,
	}
}
