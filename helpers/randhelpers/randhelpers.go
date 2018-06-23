package randhelpers

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

// GetRandom returns a copy of a random element from the array
func GetRandom(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

// GetRandomIndex returns a random index from the passed in strin
func GetRandomIndex(str string) int {
	return rand.Intn(len(str))
}

// GetRandomSliceIndex returns a random index from the passed in slice
func GetRandomSliceIndex(slice []string) int {
	return rand.Intn(len(slice))
}
