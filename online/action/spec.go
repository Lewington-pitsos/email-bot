package action

import (
	"email-bot/logger"
	"email-bot/online/browser"
)

type spec struct {
	operation
}

func (s *spec) check() bool {
	if len(s.commands) == 0 {
		logger.LoggerInterface.Println("Empty Spec: Spec passed")
		return true
	}

	return s.runChecks()
}

func (s *spec) runChecks() bool {
	for _, command := range s.commands {
		if !command(s) {
			logger.LoggerInterface.Println("Spec failed")
			return false
		}
	}
	logger.LoggerInterface.Println("Spec Passed... proceeding to interaction")
	return true
}

func NewSpec() *spec {
	return &spec{
		operation{
			commands: make([]func(*spec) bool, 0, 20),
		},
	}
}
