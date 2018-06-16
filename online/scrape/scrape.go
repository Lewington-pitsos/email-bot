package scrape

import "email-bot/online/action"

type Scrape struct {
	actions []*action.Action
}

func (s *Scrape) AddAction(action *action.Action) *Scrape {
	s.actions = append(s.actions, action)

	return s
}

func (s *Scrape) Scrape() {
	for _, action := range s.actions {
		action.Perform()
	}
}

func NewScrape() *Scrape {
	return &Scrape{
		actions: make([]*action.Action, 0, 50),
	}
}
