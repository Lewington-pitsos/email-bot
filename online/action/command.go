package action

import (
	"email-bot/online/browser"
	"email-bot/online/data"
	"email-bot/online/helpers"
	"time"
)

// +---------------------------------------------------------------------------------------+
//										INTERACTIONS
// +---------------------------------------------------------------------------------------+

func VisitPage(url string) func(*interaction) {
	return func(i *interaction) {
		i.browser.Wd.Get(url)
		time.Sleep(time.Millisecond * 3000)
	}
}

func FillField(selector string, detail *data.Detail) func(*interaction) {
	return func(i *interaction) {
		element, err := i.browser.Wd.FindElement("xpath", selector)
		helpers.CheckSafe(err)

		element.SendKeys(detail.ValueAt(i.tries))
		time.Sleep(time.Millisecond * 300)
	}
}

func Click(selector string) func(*interaction) {
	return func(i *interaction) {
		element, err := i.browser.Wd.FindElement("xpath", selector)
		helpers.CheckSafe(err)

		element.Click()
	}
}

// +---------------------------------------------------------------------------------------+
//										SPECS
// +---------------------------------------------------------------------------------------+

func CheckExists(selector string) func(*browser.Browser) bool {
	return func(browser *browser.Browser) bool {
		elements, err := browser.Wd.FindElements("xpath", selector)
		helpers.CheckSafe(err)

		return len(elements) > 0
	}
}
