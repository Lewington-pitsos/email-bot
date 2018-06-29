package action

import (
	"email-bot/datastructures"
	"email-bot/logger"
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

func (a *Action) AddToSpec(command func(*spec) bool) *Action {
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

func (a *Action) AddSelectOperation(selector string, optionSelector string, valueType string) *Action {
	a.spec.AddCommand(CheckExists(selector))
	a.interaction.AddValueCommand(SelectOption(selector+optionSelector), valueType)

	return a
}

func (a *Action) AddFillOperation(selector string, valueType string) *Action {
	a.spec.AddCommand(CheckExists(selector))
	a.interaction.AddValueCommand(FillField(selector), valueType)

	return a
}

func (a *Action) AddSubmitOperation(selector string) *Action {
	a.spec.AddCommand(CheckExists(selector))
	a.interaction.AddCommand(Click(selector))

	return a
}

func (a *Action) AddBrowser(browser *browser.Browser) *Action {
	a.interaction.addBrowser(browser)
	a.spec.addBrowser(browser)
	return a
}

func (a *Action) AddCandidateValues(candidateValues map[string]datastructures.Detail) *Action {
	logger.LoggerInterface.Println(a)
	a.interaction.addCandidateValues(candidateValues)
	return a
}

func (a *Action) Clone() *Action {
	return &Action{
		spec:        a.spec.clone(),
		interaction: a.interaction.clone(),
		Critical:    a.Critical,
	}
}

// +---------------------------------------------------------------------------------------+
//										EXPOSED FUNCTIONS
// +---------------------------------------------------------------------------------------+

func NewAction(critical bool) *Action {
	return &Action{
		spec:        NewSpec(),
		interaction: NewInteraction(),
		Critical:    critical,
	}
}
