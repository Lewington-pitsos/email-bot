package scrape

import (
	"email-bot/online/action"
	"email-bot/logger"
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
	logger.LoggerInterface.Println("Commencing Scrape")
	for i := 0; i < len(s.actions); {
		action := s.actions[i]
		result := action.Perform()
		if !result {
			logger.LoggerInterface.Println("Action failed")
			i--
		} else {
			logger.LoggerInterface.Println("Action Succeeded")
			i++
		}

		if i < 0 {
			return false
		}
		time.Sleep(time.Millisecond * 300)
	}
	
	logger.LoggerInterface.Println("Scrape Complete")
	return true
}

func NewScrape() *Scrape {
	logger.LoggerInterface.Println("Creating Scrape")
	return &Scrape{
		actions: make([]*action.Action, 0, 50),
	}
}
