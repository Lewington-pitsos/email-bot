package action

import (
	"email-bot/online/helpers"
	"time"

	"github.com/tebeka/selenium"
)

// +---------------------------------------------------------------------------------------+
//										INTERACTIONS
// +---------------------------------------------------------------------------------------+

func VisitPage(url string) func(selenium.WebDriver) {
	return func(browser selenium.WebDriver) {
		browser.Get(url)
	}
}

func FillField(selector string, value string) func(selenium.WebDriver) {
	return func(browser selenium.WebDriver) {
		element, err := browser.FindElement("xpath", selector)
		helpers.CheckSafe(err)

		element.SendKeys(value)
		time.Sleep(time.Millisecond * 300)
	}
}

func Click(selector string) func(selenium.WebDriver) {
	return func(browser selenium.WebDriver) {
		element, err := browser.FindElement("xpath", selector)
		helpers.CheckSafe(err)

		element.Click()
	}
}

// +---------------------------------------------------------------------------------------+
//										SPECS
// +---------------------------------------------------------------------------------------+

func CheckExists(selector string) func(selenium.WebDriver) bool {
	return func(browser selenium.WebDriver) bool {
		elements, err := browser.FindElements("xpath", selector)
		helpers.CheckSafe(err)

		return len(elements) > 0
	}
}
