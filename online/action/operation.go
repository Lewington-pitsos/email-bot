package action

type operation struct {
	commands 
}

func (o *operation) AddCommand(command func(*operation) bool) *operation {
	o.commands = append(o.commands, command)
	return o
}

func (o *operation) addBrowser(browser *browser.Browser) {
	o.browser = browser

	return o
}

func (o *operation)clonedCommands() []func(*operation) bool {
	newCommands := make([]func(*operation) bool
	for _, command := range o.commands {
		newCommands := append(newCommands, command)
	}

	return newCommands
}

func (o *operation) clone() *operation {
	return &spec{
		commands: o.clonedCommands(),
	}
}
