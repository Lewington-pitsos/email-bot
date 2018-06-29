package profile

type ProfileInterface struct {
	Generate()
	Values() map[string]datastructures.Detail
}