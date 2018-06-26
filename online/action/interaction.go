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

func (i *interaction) addBrowser(browser *browser.Browser) {
	i.browser = browser
}

func (i *interaction) AddCommand(command func(*interaction)) *interaction {
	i.commands = append(i.commands, command)
	return i
}

func NewInteraction() *interaction {
	return &interaction{
		commands: make([]func(*interaction), 0, 20),
	}
}
