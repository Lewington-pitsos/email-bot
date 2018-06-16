package online

import (
	"email-bot/online/scrape"
	"time"
)

func Test() {
	m := scrape.NewManager(8081)
	m.ProvisionHotmailNewAccount()
	m.Scrape()

	time.Sleep(time.Millisecond * 80000)
}
