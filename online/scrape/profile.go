package scrape

import (
	"email-bot/online/action"
)

type Profile struct {
	instructions []*action.Action
}

func (p *Profile) Instructions() []*action.Action {
	instructionClone := make([]*action.Action)
	for _, action := range i.instructions {
		instructionClone = append(instructionClone, action.Clone())
	}
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