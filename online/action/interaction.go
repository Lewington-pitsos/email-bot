package action

import "github.com/tebeka/selenium"

type interaction struct {
	tries    int
	browser  selenium.WebDriver
	operations []operation
}

func (i *interaction) run(tries int) {
	i.tries = tries
	for _, operation := range i.operations {
		opperation.Operate(i)
	}
}

func (i *interaction) AddCommand(operation *operation) *interaction {
	i.operations = append(i.operations, operation)
	return i
}

func NewInteraction(browser selenium.WebDriver) *interaction {
	return &interaction{
		tries:    0,
		browser:  browser,
		operations: make([]operation, 0, 20),
	}
}
