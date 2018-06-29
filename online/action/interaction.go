package action

import (
	"email-bot/datastructures"
	"email-bot/logger"
	"email-bot/online/browser"
)

type interaction struct {
	browser         *browser.Browser
	commands        []func(*interaction)
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

func (i *interaction) CurrentValue() string {
	value := i.candidateValues[i.valueIndex].RandomValue()
	i.valueIndex++
	return value
}

func (i *interaction) addBrowser(browser *browser.Browser) {
	i.browser = browser
}

func (i *interaction) AddValueCommand(command func(*interaction), valueType string) *interaction {
	i.AddCommand(command)
	i.addValueType(valueType)
	return i
}

func (i *interaction) AddCommand(command func(*interaction)) *interaction {
	i.commands = append(i.commands, command)
	return i
}

func (i *interaction) addValueType(valueType string) {
	i.valueTypes = append(i.valueTypes, valueType)
}

func (i *interaction) clonedValueTypes() []string {
	newValueTypes := make([]string, 0, 20)
	for _, valueType := range i.valueTypes {
		newValueTypes = append(newValueTypes, valueType)
	}

	return newValueTypes
}

func (i *interaction) clonedCommands() []func(*interaction) {
	newCommands := make([]func(*interaction), 0, 20)
	for _, command := range i.commands {
		newCommands = append(newCommands, command)
	}

	return newCommands
}

func (i *interaction) clone() *interaction {
	return &interaction{
		commands:        i.clonedCommands(),
		candidateValues: make([]datastructures.Detail, 0, 20),
		valueTypes:      i.clonedValueTypes(),
		valueIndex:      0,
	}
}

func (i *interaction) addCandidateValues(candidateValues map[string]datastructures.Detail) *interaction {
	for _, valueType := range i.valueTypes {
		i.candidateValues = append(i.candidateValues, candidateValues[valueType])
	}
	return i
}

func NewInteraction() *interaction {
	return &interaction{
		commands:        make([]func(*interaction), 0, 20),
		candidateValues: make([]datastructures.Detail, 0, 20),
		valueTypes:      make([]string, 0, 20),
		valueIndex:      0,
	}
}
