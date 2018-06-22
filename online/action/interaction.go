package action

import (
	"email-bot/online/browser"
	"email-bot/logger"
	"fmt"
)

type interaction struct {
	tries    int
	browser  *browser.Browser
	commands []func(*interaction)
}

func (i *interaction) run() {
	i.tries++
	fmt.Println(i.commands)
	logger.LoggerInterface.Println("Running interaction, try number:", i.tries)
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
		tries:    0,
		browser:  browser,
		commands: make([]func(*interaction), 0, 20),
	}
}
