package action

type operation struct {
	Values []string
	Tries int
	Command func(interface{})
}

func(o *operation) Operate()  {
	o.Command()
}

func NewOperation(values []string, command func(interface{})) * operation {
	return &operation {
		Values: values,
		Tries: 0, 
		Command: command,
	}
}