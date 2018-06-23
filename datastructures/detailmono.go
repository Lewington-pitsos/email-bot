package datastructures

type DetailMono struct {
	value string
}

func (d *DetailMono) RandomValue() string {
	return d.value
}

func (d *DetailMono) CurrentValue() string {
	return d.value
}

func NewDetaiMono(value string) *DetailMono {
	return &DetailMono{
		value: value,
	}
}
