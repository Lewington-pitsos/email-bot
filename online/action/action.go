package action

import (
	"fmt"

	"github.com/tebeka/selenium"
)

type Action struct {
	Tries		int
	spec        *spec
	interaction *interaction
}

func (a *Action) Perform() bool {
	fmt.Println(a)
	if a.spec.check() {
		a.interaction.run(a.Tries)
		a.Tries++
		fmt.Println("interaction performed")
		return true
	}

	fmt.Println("spec failed")

	return false
}

func (a *Action) AddToInteraction(command func(*interaction)) *Action {
	a.interaction.AddCommand(command)

	return a
}

func (a *Action) AddToSpec(command func(selenium.WebDriver) bool) *Action {
	a.spec.AddCommand(command)

	return a
}

func (a *Action) AddFillOperation(selector string, value []string) *Action {
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
		Tries: 		0,
		spec:        NewSpec(browser),
		interaction: NewInteraction(browser),
	}
}
