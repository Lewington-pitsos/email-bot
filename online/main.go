package online

import (
	"email-bot/online/action"
	"email-bot/online/crawler"
	"email-bot/online/scrape"
	"fmt"
	"time"
)

func Test() {
	wd := crawler.NewWebDriver(8081)

	s := action.NewSpec(wd)
	i := action.NewInteraction(wd)

	command := action.VisitPage("https://google.com")
	i.AddCommand(command)

	a := action.NewAction(s, i)

	scrape := scrape.NewScrape()
	scrape.AddAction(a)

	fmt.Println(scrape)
	scrape.Scrape()

	time.Sleep(time.Millisecond * 20000)
	wd.Quit()
}
