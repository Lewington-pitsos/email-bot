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


func (s *Scrape) AddActions(actions []*action.Action) *Scrape {
	for _, action := range actions {
		s.actions = append(s.actions, action)
	}

	return s
}

func (s *Scrape) Scrape() bool {
	for i := 0; i < len(s.actions); {
		action := s.actions[i]
		result := action.Perform()
		fmt.Println(result)
		if !result {
			i--
		} else {
			i++
		}

		if i < 0 {
			return false
		}
		time.Sleep(time.Millisecond * 300)
	}

	return true
}

func NewScrape() *Scrape {
	return &Scrape{
		actions: make([]*action.Action, 0, 50),
	}
}
