package action

import (
	"email-bot/logger"
	"email-bot/online/browser"
	"fmt"
)

type interaction struct {
	browser  *browser.Browser
	commands []func(*interaction)
}

func (i *interaction) run() {
	fmt.Println(i.commands)
	logger.LoggerInterface.Println("Running interaction, value number:")
	for _, command := range i.commands {
		command(i)
	}
}

func (i *interaction) AddCommand(command func(*interaction)) *interaction {
	i.commands = append(i.commands, command)
	return i
}

func NewInteraction(browser *browser.Browser) *interaction {
	return &interaction{
		browser:  browser,
		commands: make([]func(*interaction), 0, 20),
	}
}
