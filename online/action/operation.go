package action

import "github.com/tebeka/selenium"

type operation struct {
	browser  selenium.WebDriver
	commands []func()
}

func (o *operation) addCommand(command func()) *operation {
	o.commands = append(o.commands, command)
	return o
}
