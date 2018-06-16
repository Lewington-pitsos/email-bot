package scrape

import (
	"email-bot/online/action"
	"fmt"
	"time"
)

type Scrape struct {
	actions []*action.Action
}

func (s *Scrape) AddActions(actions ...*action.Action) *Scrape {
	for _, action := range actions {
		s.actions = append(s.actions, action)
	}

	return s
}

func (s *Scrape) Scrape() {
	for _, action := range s.actions {
		time.Sleep(time.Millisecond * 2000)
		fmt.Println(action)
		action.Perform()
	}
}

func NewScrape() *Scrape {
	return &Scrape{
		actions: make([]*action.Action, 0, 50),
	}
}
