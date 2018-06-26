package scrape

import (
	"email-bot/online/action"

	"github.com/getlantern/deepcopy"
)

type Profile struct {
	instructions []*action.Action
}

func (p *Profile) Instructions() []*action.Action {
	instructionClone := make([]*action.Action, 0, 40)
	deepcopy.Copy(instructionClone, p.instructions)
	return instructionClone
}

func (p *Profile) AddAction(newAction *action.Action) {
	p.instructions = append(p.instructions, newAction)
}

func NewProfile() *Profile {
	return &Profile{
		instructions: make([]*action.Action, 0, 40),
	}
}
