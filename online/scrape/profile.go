package scrape

import (
	"email-bot/online/action"
)

type Profile struct {
	instructions []*action.Action
}

func (p *Profile) Instructions() []*action.Action {
	instructionClone := make([]*action.Action, 0, 20)
	for _, action := range p.instructions {
		instructionClone = append(instructionClone, action.Clone())
	}
	return instructionClone
}

func (p *Profile) AddAction(critical bool) *action.Action {
	newAction := action.NewAction(critical)
	p.instructions = append(p.instructions, newAction)
	return newAction
}

func NewProfile() *Profile {
	return &Profile{
		instructions: make([]*action.Action, 0, 40),
	}
}
