package action

import (
	"email-bot/logger"
	"email-bot/online/browser"
	"email-bot/online/data"
	"email-bot/online/helpers"
)

// +---------------------------------------------------------------------------------------+
//										INTERACTIONS
// +---------------------------------------------------------------------------------------+

func VisitPage(url string) func(*interaction) {
	return func(i *interaction) {
		logger.LoggerInterface.Println("Visiting:", url)
		i.browser.Wd.Get(url)
	}
}

func FillField(selector string, detail *data.Detail) func(*interaction) {
	return func(i *interaction) {
		element, err := i.browser.Wd.FindElement("xpath", selector)
		helpers.CheckSafe(err)

		value := detail.ValueAt(i.tries)
		logger.LoggerInterface.Println("Filling element:", selector, "with:", value)
		element.Clear()

		element.SendKeys(value)
	}
}

func Click(selector string) func(*interaction) {
	return func(i *interaction) {
		logger.LoggerInterface.Println("Clicking element:", selector)
		element, err := i.browser.Wd.FindElement("xpath", selector)
		helpers.CheckSafe(err)

		element.Click()
	}
}

func Wait(wait int) func(*interaction) {
	return func(i *interaction) {
		logger.LoggerInterface.Println("Waiting:", wait, "miliseconds")
		helpers.Wait(wait)
	}
}

// +---------------------------------------------------------------------------------------+
//										SPECS
// +---------------------------------------------------------------------------------------+

func CheckExists(selector string) func(*browser.Browser) bool {
	return func(browser *browser.Browser) bool {
		logger.LoggerInterface.Println("Checking for element:", selector)
		elements, err := browser.Wd.FindElements("xpath", selector)
		helpers.CheckSafe(err)

		return len(elements) > 0
	}
}
