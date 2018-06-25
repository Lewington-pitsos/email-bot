package action

import (
	"email-bot/datastructures"
	"email-bot/online/browser"
)

// +---------------------------------------------------------------------------------------+
//										Action STRUCT
// +---------------------------------------------------------------------------------------+

type Action struct {
	spec        *spec
	interaction *interaction
	Critical    bool
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

func (a *Action) AddVisit(url string) *Action {
	a.interaction.AddCommand(VisitPage(url))

	return a
}

func (a *Action) AddSelectOperation(selector string, optionSelector string, detail datastructures.Detail) *Action {
	a.spec.AddCommand(CheckExists(selector))
	a.interaction.AddCommand(SelectOption(selector+optionSelector, detail))

	return a
}

func (a *Action) AddFillOperation(selector string, detail datastructures.Detail) *Action {
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

func NewAction(browser *browser.Browser, critical bool) *Action {
	return &Action{
		spec:        NewSpec(browser),
		interaction: NewInteraction(browser),
		Critical:    critical,
	}
}
