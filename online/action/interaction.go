package action

import "github.com/tebeka/selenium"

type interaction struct {
	browser  selenium.WebDriver
	commands []func(selenium.WebDriver)
}

func (i *interaction) run() {
	for _, command := range i.commands {
		command(i.browser)
	}
}

func (i *interaction) AddCommand(command func(selenium.WebDriver)) *interaction {
	i.commands = append(i.commands, command)
	return i
}

func NewInteraction(browser selenium.WebDriver) *interaction {
	return &interaction{
		browser:  browser,
		commands: make([]func(selenium.WebDriver), 0, 20),
	}
}
