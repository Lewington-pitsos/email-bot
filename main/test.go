package main

import (
	"email-bot/helpers/generalhelpers"
	"email-bot/online/browser"
	"fmt"
)

func main() {
	b := browser.NewBrowser(3334)
	b.Wd.Get("https://beteasy.com.au/?fbclid=IwAR1nHIJvuIdCYvXTS5Pq-Z8CXUGRrPzUFOXdAjFZv5AFQ57tZVqO9PUdGds")
	x, err := b.Wd.FindElements("xpath", "//span[text()='Australia/NZ' and contains(@class, 'hdo-location__type')]/parent::div[contains(@class, 'hdo-location')]//a[contains(@class, 'hdo-event')]")
	generalhelpers.Check(err)

	for _, i := range x {
		j, err := i.GetAttribute("href")
		generalhelpers.Check(err)
		fmt.Println(j)
	}
}
