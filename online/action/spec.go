package action

import (
	"email-bot/logger"
	"email-bot/online/browser"
)

type spec struct {
	browser  *browser.Browser
	commands []func(*browser.Browser) bool
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
		if !command(s.browser) {
			logger.LoggerInterface.Println("Spec failed")
			return false
		}
	}
	logger.LoggerInterface.Println("Spec Passed... proceeding to interaction")
	return true
}

func NewSpec(br *browser.Browser) *spec {
	return &spec{
		browser:  br,
		commands: make([]func(*browser.Browser) bool, 0, 20),
	}
}

func (s *spec) AddCommand(command func(*browser.Browser) bool) *spec {
	s.commands = append(s.commands, command)
	return s
}
