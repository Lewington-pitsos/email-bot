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
		AddVisit("https://www.ladbrokes.com.au/racing/").
		AddWait(300)

	b.AddAction(true).
		AddExtractOperation("//div[@id='racesToday']/div[contains(@class, 'fullbox')][1]//td[contains(@class, 'odds')]//a[contains(@class, 'subpage')]", "href", channel)

	b.Scrape()

	time.Sleep(time.Millisecond * 80000)

	// channel := make(chan *datastructures.Signal)
	// r := relay.NewRelay(channel)
	// r.Listen()

	// b := emailbot.NewBot(3334)

	// b.AddAction(true).
	// 	AddVisit("https://beteasy.com.au/?fbclid=IwAR1nHIJvuIdCYvXTS5Pq-Z8CXUGRrPzUFOXdAjFZv5AFQ57tZVqO9PUdGds").
	// 	AddWait(300)

	// b.AddAction(true).
	// 	AddExtractOperation("visit-link", "//span[text()='Australia/NZ' and contains(@class, 'hdo-location__type')]/parent::div[contains(@class, 'hdo-location')]//a[contains(@class, 'hdo-event')]", "href", channel)

	// b.Scrape()

	// time.Sleep(time.Millisecond * 80000)
}
