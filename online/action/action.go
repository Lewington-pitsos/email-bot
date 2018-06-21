package action

import (
	"email-bot/online/browser"
	"email-bot/online/data"
	"fmt"
)

type Action struct {
	spec        *spec
	interaction *interaction
}

func (a *Action) Perform() bool {
	if a.spec.check() {
		a.interaction.run()
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

func (a *Action) AddToSpec(command func(*browser.Browser) bool) *Action {
	a.spec.AddCommand(command)

	return a
}

func (a *Action) AddWait(wait int) *Action {
	a.interaction.AddCommand(Wait(wait))

	return a
}

func (a *Action) AddFillOperation(selector string, detail *data.Detail) *Action {
	a.spec.AddCommand(CheckExists(selector))
	a.interaction.AddCommand(FillField(selector, detail))

	return a
}

func (a *Action) AddSubmitOperation(selector string) *Action {
	a.spec.AddCommand(CheckExists(selector))
	a.interaction.AddCommand(Click(selector))

	return a
}

func NewAction(browser *browser.Browser) *Action {
	return &Action{
		spec:        NewSpec(browser),
		interaction: NewInteraction(browser),
	}
}
