package profile

type Profile struct {
	Name   string
	fields []*field
}

func (p *Profile) addField(d *field) *Profile {
	p.fields = append(p.fields, d)

	return p
}

func NewProfile(name string) *Profile {
	return &Profile{
		Name:   name,
		fields: make([]*field, 0, 50),
	}
}
