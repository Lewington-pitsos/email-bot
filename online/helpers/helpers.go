package helpers

import "time"

func CheckSafe(err error) {
	if err != nil {
		panic(err)
	}
}

func Wait(milliseconds int) {
	if milliseconds > 0 {
		time.Sleep(time.Millisecond * time.Duration(milliseconds))
	}
}
