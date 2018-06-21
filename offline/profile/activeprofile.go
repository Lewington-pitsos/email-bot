package profile

import (
	"email-bot/offline/files"
	"fmt"
)

type ActiveProfile struct {
	Values map[string]string
	Name   string
}

func (a *ActiveProfile) Populate() {
	loader := files.NewManager()
	data := loader.ProfileData(a.Name)
	fmt.Println(data)
}

func NewActiveProfile(name string) *ActiveProfile {
	return &ActiveProfile{
		Values: make(map[string]string),
		Name:   name,
	}
}
