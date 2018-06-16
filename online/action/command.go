package action

import "github.com/tebeka/selenium"

func VisitPage(url string, browser selenium.WebDriver) func() {
	return func() {
		browser.Get(url)
	}
}
