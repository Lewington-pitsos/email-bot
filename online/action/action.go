package action

import (
	"fmt"
)

type Action struct {
	expectation *spec
	interaction *interaction
}

func (a *Action) Perform() bool {
	fmt.Println(a)
	if a.expectation.check() {
		a.interaction.run()
		return true
	}

	return false
}

func NewAction(expectation *spec, interaction *interaction) *Action {
	return &Action{
		expectation: expectation,
		interaction: interaction,
	}
}
