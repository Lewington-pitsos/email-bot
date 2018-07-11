package profile

import (
	"email-bot/datastructures"
	"email-bot/logger"
	"email-bot/offline/chunk"
	"email-bot/offline/generator"
	"email-bot/offline/valuebank"
)

type DataProfile struct {
	values    map[string]datastructures.Detail
	fields    []*Field
	generator *generator.ValueGenerator
	Bank      *valuebank.Bank
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

func (p *DataProfile) Values() map[string]datastructures.Detail {
	newMap := make(map[string]datastructures.Detail)
	for key, value := range p.values {
		newMap[key] = value
	}

	return newMap
}

func (p *DataProfile) AddToBank(fileName string, vaultName string) *DataProfile {
	p.generator.AddToBank(fileName, vaultName)

	return p
}

func (p *DataProfile) WithModifiedChunk(mode string, source string, modBank string) *DataProfile {
	modChunk := chunk.NewChunk(mode, source).SetModification(modBank)
	field := p.fields[len(p.fields)-1]
	field.Format = append(field.Format, modChunk)

	return p
}

func (p *DataProfile) Generate() ProfileInterface {
	logger.LoggerInterface.Println("Generating DataProfile data")
	for _, field := range p.fields {
		p.values[field.Name] = p.generator.Generate(field.Format, field.ValueNumber)
	}

	return p
}

func (d *DataProfile) Saveable() bool {
	return true
}

func NewDataProfile() *DataProfile {
	generator := generator.NewValueGenerator()
	logger.LoggerInterface.Println("Creating DataProfile")
	values := make(map[string]datastructures.Detail)
	generator.SetValues(values)

	return &DataProfile{
		values:    values,
		fields:    make([]*Field, 0, 50),
		generator: generator,
	}
}
