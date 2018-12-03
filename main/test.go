package main

import (
	"email-bot/datastructures"
	"email-bot/emailbot"
	"email-bot/relay"
	"time"
)

func main() {
	channel := make(chan *datastructures.Signal)
	r := relay.NewRelay(channel)
	r.Listen()

	b := emailbot.NewBot(3334)

	b.AddAction(true).
		AddVisit("https://beteasy.com.au/?fbclid=IwAR1nHIJvuIdCYvXTS5Pq-Z8CXUGRrPzUFOXdAjFZv5AFQ57tZVqO9PUdGds").
		AddWait(300)

	b.AddAction(true).
		AddExtractOperation("visit-link", "//span[text()='Australia/NZ' and contains(@class, 'hdo-location__type')]/parent::div[contains(@class, 'hdo-location')]//a[contains(@class, 'hdo-event')]", "href", channel)

	b.Scrape()

	time.Sleep(time.Millisecond * 80000)

	// b := browser.NewBrowser(3334)
	// b.Wd.Get("https://beteasy.com.au/?fbclid=IwAR1nHIJvuIdCYvXTS5Pq-Z8CXUGRrPzUFOXdAjFZv5AFQ57tZVqO9PUdGds")
	// x, err := b.Wd.FindElements("xpath", "//span[text()='Australia/NZ' and contains(@class, 'hdo-location__type')]/parent::div[contains(@class, 'hdo-location')]//a[contains(@class, 'hdo-event')]")
	// generalhelpers.Check(err)

	// for _, i := range x {
	// 	j, err := i.GetAttribute("href")
	// 	generalhelpers.Check(err)
	// 	fmt.Println(j)
	// }
}
