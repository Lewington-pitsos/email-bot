package action

import (
	"email-bot/online/browser"
	"fmt"
)

type spec struct {
	browser  *browser.Browser
	commands []func(*browser.Browser) bool
}

func (s *spec) check() bool {
	if len(s.commands) == 0 {
		return true
	}

	fmt.Println("Spec Passed")

	return s.runChecks()
}

func (s *spec) runChecks() bool {
	for _, command := range s.commands {
		if !command(s.browser) {
			return false
		}
	}

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
