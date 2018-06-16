package online

import (
	"email-bot/online/crawler"
	"time"
)

func Test() {
	wd := crawler.NewWebDriver(8081)

	wd.Get("http://google.com")

	time.Sleep(time.Millisecond * 20000)
	wd.Quit()
}
