package helpers

import (
	"math/rand"
	"time"
)

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRandom returns a copy of a random element from the array
func GetRandom(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

// GetRandomIndex returns a random index from the passed in slice
func GetRandomIndex(str string) int {
	return rand.Intn(len(str))
}

func Check(err error) {
	if	err != nil {
		panic(err)
	}
}