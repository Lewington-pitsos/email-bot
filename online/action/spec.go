package action

import (
	"fmt"

	"github.com/tebeka/selenium"
)

type spec struct {
	browser  selenium.WebDriver
	commands []func() bool
}

func (s *spec) check() bool {
	if len(s.commands) == 0 {
		return true
	}

	fmt.Println("loool")

	return s.runChecks()
}

func (s *spec) runChecks() bool {
	for _, command := range s.commands {
		if command() {
			return false
		}
	}

	return true
}

func (s *spec) addCommand(command func() bool) *spec {
	s.commands = append(s.commands, command)
	return s
}

func NewSpec(browser selenium.WebDriver) *spec {
	return &spec{
		browser:  browser,
		commands: make([]func() bool, 0, 20),
	}
}
