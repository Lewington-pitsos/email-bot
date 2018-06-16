package action

import (
	"email-bot/online/helpers"
	"time"

	"github.com/tebeka/selenium"
)

// +---------------------------------------------------------------------------------------+
//										INTERACTIONS
// +---------------------------------------------------------------------------------------+

func VisitPage(url string) func(*interaction) {
	return func(i *interaction) {
		i.browser.Get(url)
		time.Sleep(time.Millisecond * 3000)
	}
}

func FillField(selector string, value []string) func(*interaction) {
	return func(i *interaction) {
		element, err := i.browser.FindElement("xpath", selector)
		helpers.CheckSafe(err)

		element.SendKeys(value[i.tries])
		time.Sleep(time.Millisecond * 300)
	}
}

func Click(selector string) func(*interaction) {
	return func(i *interaction) {
		element, err := i.browser.FindElement("xpath", selector)
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
