package action

import (
	"email-bot/helpers/generalhelpers"
	"email-bot/logger"
	"fmt"
)

// +---------------------------------------------------------------------------------------+
//										INTERACTIONS
// +---------------------------------------------------------------------------------------+

func VisitPage(url string) func(*operationInterface) {
	return func(o *operationInterface) {
		logger.LoggerInterface.Println("Visiting:", url)
		o.browser.Wd.Get(url)
	}
}

func FillField(selector string) func(*interaction) {
	return func(i *interaction) {
		element, err := i.browser.Wd.FindElement("xpath", selector)
		generalhelper.Check(err)

		value := i.CurrentValue()
		logger.LoggerInterface.Println("Filling element:", selector, "with:", value)
		element.Clear()

		element.SendKeys(value)
	}
}

func SelectOption(selector string) func(*interaction) {
	return func(i *interaction) {
		value := i.CurrentValue()
		logger.LoggerInterface.Println("Setting Selection:", selector, "as:", value)

		optionSelector := fmt.Sprintf(selector, value)

		logger.LoggerInterface.Println(optionSelector)

		element, err := i.browser.Wd.FindElement("xpath", optionSelector)
		generalhelpers.Check(err)

		element.Click()
	}
}

func Click(selector string) func(*operationInterface) {
	return func(o *operationInterface) {
		logger.LoggerInterface.Println("Clicking element:", selector)
		element, err := o.browser.Wd.FindElement("xpath", selector)
		generalhelpers.Check(err)

		element.Click()
	}
}

func Wait(wait int) func(*operationInterface) {
	return func(o *operationInterface) {
		logger.LoggerInterface.Println("Waiting:", wait, "miliseconds")
		generalhelpers.Wait(wait)
	}
}

// +---------------------------------------------------------------------------------------+
//										SPECS
// +---------------------------------------------------------------------------------------+

func CheckExists(selector string) func(*operationInterface) bool {
	return func(o *operationInterface) bool {
		logger.LoggerInterface.Println("Checking for element:", selector)
		elements, err := o.browser.Wd.FindElements("xpath", selector)
		generalhelpers.Check(err)

		return len(elements) > 0
	}
}
