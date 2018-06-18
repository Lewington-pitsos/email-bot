package data

type Detail struct {
	values  []string
	current int
}

func (d *Detail) ValueAt(index int) string {
	d.current = index
	return d.values[index]
}

func (d *Detail) CurrentValue() string {
	return d.values[d.current]
}

func NewDetail(values []string) *Detail {
	return &Detail{
		values:  values,
		current: 0,
	}
}
