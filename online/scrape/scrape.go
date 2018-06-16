package scrape

import (
	"email-bot/online/action"
	"fmt"
	"time"
)

type Scrape struct {
	actions []*action.Action
}

func (s *Scrape) AddAction(action *action.Action) *Scrape {
	s.actions = append(s.actions, action)

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
