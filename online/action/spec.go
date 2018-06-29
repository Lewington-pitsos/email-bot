package action

import (
	"email-bot/logger"
	"email-bot/online/browser"
)

type spec struct {
	browser  *browser.Browser
	commands []func(*spec) bool
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

func (s *spec) addBrowser(browser *browser.Browser) {
	s.browser = browser
}

func (s *spec) clonedCommands() []func(*spec) bool {
	newCommands := make([]func(*spec) bool, 0, 20)
	for _, command := range s.commands {
		newCommands = append(newCommands, command)
	}

	return newCommands
}

func (s *spec) clone() *spec {
	return &spec{
		commands: s.clonedCommands(),
	}
}

func (s *spec) AddCommand(command func(*spec) bool) *spec {
	s.commands = append(s.commands, command)
	return s
}

func NewSpec() *spec {
	return &spec{
		commands: make([]func(*spec) bool, 0, 20),
	}
}
