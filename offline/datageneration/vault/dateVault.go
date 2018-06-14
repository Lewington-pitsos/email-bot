package vault

import (
	"math/rand"
	"time"
)

var StartTime int = -410227200
var endTime int = 915148800

type DateVault struct {
}

func (d *DateVault) GiveValue() string {
	var unixTime = int64(rand.Intn(endTime-StartTime) + StartTime)
	return time.Unix(unixTime, 0).Format("02/01/2006")
}

func NewDateVault() *DateVault {
	return &DateVault{}
}
