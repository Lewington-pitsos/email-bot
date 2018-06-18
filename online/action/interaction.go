package action

import (
	"email-bot/online/browser"
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
