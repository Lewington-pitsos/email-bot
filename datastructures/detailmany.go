package datastructures

import "email-bot/helpers/randhelpers"

type DetailMany struct {
	values  []string
	current int
}

func (d *DetailMany) RandomValue() string {
	return d.valueAt(randhelpers.GetRandomSliceIndex(d.values))
}

func (d *DetailMany) valueAt(index int) string {
	d.current = index
	return d.values[index]
}

func (d *DetailMany) CurrentValue() string {
	return d.values[d.current]
}

func NewDetaiMany(values []string) *DetailMany {
	return &DetailMany{
		values:  values,
		current: 0,
	}
}
