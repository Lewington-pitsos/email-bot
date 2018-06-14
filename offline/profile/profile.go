package profile

type Profile struct {
	Name   string
	fields []*field
}

func (p *Profile) AddField(d *field) *Profile {
	p.fields = append(p.fields, d)

	return p
}

func (p *Profile) Profile() *map[string][]string {
	profile := make(map[string][]string)
	for _, field := range p.fields {
		profile[field.Name] = field.Values
	}

	return &profile
}

func NewProfile(name string) *Profile {
	return &Profile{
		Name:   name,
		fields: make([]*field, 0, 50),
	}
}
