package action

import "github.com/tebeka/selenium"

func VisitPage(url string) func(selenium.WebDriver) {
	return func(browser selenium.WebDriver) {
		browser.Get(url)
	}
}
