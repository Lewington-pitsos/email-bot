package action

import (
	"email-bot/helpers/generalhelpers"
	"email-bot/logger"
	"fmt"
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

func FillField(selector string) func(*interaction) {
	return func(i *interaction) {
		element, err := i.browser.Wd.FindElement("xpath", selector)
		generalhelpers.Check(err)

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

func ElementCount(spec *spec, selector string) int {
	elements, err := spec.browser.Wd.FindElements("xpath", selector)
	generalhelpers.Check(err)

	return len(elements)
}

func CheckExists(selector string) func(*spec) bool {
	return func(s *spec) bool {
		logger.LoggerInterface.Println("Checking for element:", selector)

		return ElementCount(s, selector) > 0
	}
}

func CheckDoesntExist(selector string) func(*spec) bool {
	return func(s *spec) bool {
		logger.LoggerInterface.Println("Confirming missing element:", selector)

		return ElementCount(s, selector) <= 1
	}
}
