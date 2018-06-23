package action

import (
	"email-bot/helpers/generalhelpers"
	"email-bot/logger"
	"email-bot/online/browser"
	"email-bot/online/data"
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
		generalhelpers.Check(err)

		value := detail.RandomValue()
		logger.LoggerInterface.Println("Filling element:", selector, "with:", value)
		element.Clear()

		element.SendKeys(value)
	}
}

func Click(selector string) func(*interaction) {
	return func(i *interaction) {
		logger.LoggerInterface.Println("Clicking element:", selector)
		element, err := i.browser.Wd.FindElement("xpath", selector)
		generalhelpers.Check(err)

		element.Click()
	}
}

func Wait(wait int) func(*interaction) {
	return func(i *interaction) {
		logger.LoggerInterface.Println("Waiting:", wait, "miliseconds")
		generalhelpers.Wait(wait)
	}
}

// +---------------------------------------------------------------------------------------+
//										SPECS
// +---------------------------------------------------------------------------------------+

func CheckExists(selector string) func(*browser.Browser) bool {
	return func(browser *browser.Browser) bool {
		logger.LoggerInterface.Println("Checking for element:", selector)
		elements, err := browser.Wd.FindElements("xpath", selector)
		generalhelpers.Check(err)

		return len(elements) > 0
	}
}
