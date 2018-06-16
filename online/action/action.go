package action

import (
	"fmt"

	"github.com/tebeka/selenium"
)

type Action struct {
	spec        *spec
	interaction *interaction
}

func (a *Action) Perform() bool {
	fmt.Println(a)
	if a.spec.check() {
		a.interaction.run()
		return true
	}

	return false
}

func (a *Action) AddToInteraction(command func(selenium.WebDriver)) *Action {
	a.interaction.AddCommand(command)

	return a
}

func (a *Action) AddToSpec(command func(selenium.WebDriver) bool) *Action {
	a.spec.AddCommand(command)

	return a
}

func (a *Action) AddFillOperation(selector string, value string) *Action {
	a.spec.AddCommand(CheckExists(selector))
	a.interaction.AddCommand(FillField(selector, value))

	return a
}

func (a *Action) AddSubmitOperation(selector string) *Action {
	a.spec.AddCommand(CheckExists(selector))
	a.interaction.AddCommand(Click(selector))

	return a
}

func NewAction(browser selenium.WebDriver) *Action {
	return &Action{
		spec:        NewSpec(browser),
		interaction: NewInteraction(browser),
	}
}
