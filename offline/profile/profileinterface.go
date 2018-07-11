package profile

import "email-bot/datastructures"

type ProfileInterface interface {
	Generate() ProfileInterface
	Values() map[string]datastructures.Detail
	Saveable() bool
}
