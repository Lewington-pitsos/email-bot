package action

import (
	"email-bot/online/browser"
	"email-bot/online/data"
	"fmt"
)

// +---------------------------------------------------------------------------------------+
//										Action STRUCT
// +---------------------------------------------------------------------------------------+


type Action struct {
	spec        *spec
	interaction *interaction
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (a *Action) Perform() bool {
	if a.spec.check() {
		a.interaction.run()
		return true
	}

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

// +---------------------------------------------------------------------------------------+
//										EXPOSED FUNCTIONS
// +---------------------------------------------------------------------------------------+

func NewAction(browser *browser.Browser) *Action {
	return &Action{
		spec:        NewSpec(browser),
		interaction: NewInteraction(browser),
	}
}
