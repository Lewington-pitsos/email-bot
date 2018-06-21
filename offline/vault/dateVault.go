package vault

import (
	"math/rand"
	"time"
)

const StartTime int = -410227200
const endTime int = 915148800

type DateVault struct {
	format string
}

func (d *DateVault) GiveValue() string {
	var unixTime = int64(rand.Intn(endTime-StartTime) + StartTime)
	return time.Unix(unixTime, 0).Format(d.format)
}

func NewDateVault(format string) *DateVault {
	return &DateVault{
		format: format,
	}
}
