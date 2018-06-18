package vault

import (
	"math/rand"
)

const length int = 20

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")

type PassVault struct {
}

func (p *PassVault) GiveValue() string {
	return p.RandomString()
}

func (p *PassVault) RandomString() string {
	runes := make([]rune, length)
	for i := range runes {
		runes[i] = chars[rand.Intn(len(chars))]
	}

	return string(runes)
}

func NewPassVault() *PassVault {
	return &PassVault{}
}
