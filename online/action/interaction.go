package action

import (
	"email-bot/datastructures"
	"email-bot/logger"
	"email-bot/online/browser"
)

type interaction struct {
	operation
	valueTypes      []string
	candidateValues []datastructures.Detail
	valueIndex      int
}

func (i *interaction) run() {
	logger.LoggerInterface.Println("Running interaction, value number:")
	i.valueIndex = 0
	for _, command := range i.commands {
		command(i)
	}
}

func (i *interaction) addValueType(valueType string) {
	i.valueTypes = append(i.valueTypes, valueType)
}

func (i *interaction)clonedValueTypes() []string {
	newValueTypes := make([]func(*interaction)
	for _, valueType := range i.valueTypes {
		newValueTypes := append(newValueTypes, valueType)
	}

	return newValueTypes
}

func (i *interaction) addCandidateValues(candidateValues map[string]datastructures.Detail) *interaction {
	for _, valueType := range i.valueTypes {
		i.candidateValues = append(i.candidateValues, candidateValues[valueType])
	}
	return i
}

func (i *interaction) CurrentValue() string {
	value := i.candidateValues[i.valueIndex].RandomValue()
	i.valueIndex++
	return value
}

func (i *interaction) AddValueCommand(command func(*interaction), valueType string) *interaction {
	i.AddCommand(command)
	i.addValueType(valueType)
	return i
}

func NewInteraction() *interaction {
	return &interaction{
		candidateValues: make([]datastructures.Detail, 0, 20),
		valueTypes:      make([]string, 0, 20),
		valueIndex:      0,
		operation{
			commands: make([]func(*spec) bool, 0, 20),
		},
	}
}
