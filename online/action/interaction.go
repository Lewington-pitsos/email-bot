package action

type interaction struct {
	operation
}

func (i *interaction) interact() {
	for _, command := range i.commands {
		command()
	}
}
