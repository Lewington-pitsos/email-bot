package scrape

import (
	"email-bot/logger"
	"email-bot/online/action"
)

// +-------------------------------------------------------------------------------------+
// 									Vault STRUCT
// +-------------------------------------------------------------------------------------+

const maxFails int = 30

type Scrape struct {
	actions            []*action.Action
	fails              int
	continueScrape     bool
	currentActionIndex int
	Success            bool
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (s *Scrape) currentAction() *action.Action {
	return s.actions[s.currentActionIndex]
}

func (s *Scrape) moveToNextAction() {
	s.currentActionIndex++
	if s.currentActionIndex >= len(s.actions) {
		s.continueScrape = false
		s.Success = true
		logger.LoggerInterface.Println("Scrape Completed Successfully")
	}
}

func (s *Scrape) returnToPreviousAction() {
	s.currentActionIndex--
}

func (s *Scrape) handleFailure(failedAction *action.Action) {
	if failedAction.Critical {
		logger.LoggerInterface.Println("\n\n>>>>>>>>>>>>>>>> Scrape Failed: Critical Action Failed <<<<<<<<<<<<<<<<\n\n")
		s.continueScrape = false
	} else if s.fails <= maxFails {
		logger.LoggerInterface.Println("\n\n>>>>>>>>>>>>>>>> Scrape Failed: Too Many Failures <<<<<<<<<<<<<<<<\n\n")
		s.continueScrape = false
	} else {
		s.returnToPreviousAction()
	}
}

func (s *Scrape) determineNextAction(perfomedAction *action.Action, success bool) {
	if success {
		logger.LoggerInterface.Println("Action Succeeded\n")
		s.moveToNextAction()
	} else {
		logger.LoggerInterface.Println("Action failed\n")
		s.handleFailure(perfomedAction)
	}
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

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

func (s *Scrape) Scrape() {
	logger.LoggerInterface.Println("============= Commencing Scrape =============\n\n")
	for s.continueScrape {
		action := s.currentAction()
		success := action.Perform()
		s.determineNextAction(action, success)
	}
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

func NewScrape() *Scrape {
	logger.LoggerInterface.Println("\n\n============= Creating Scrape =============")
	return &Scrape{
		actions:            make([]*action.Action, 0, 50),
		fails:              0,
		continueScrape:     true,
		currentActionIndex: 0,
		Success:            false,
	}
}
