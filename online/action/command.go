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

func SelectOption(selector string, option string) func(*interaction) {
	return func(i *interaction) {
		element, err := i.browser.Wd.FindElement("xpath", selector)
		generalhelpers.Check(err)
		element.Click()

		generalhelpers.Wait(300)

		value := i.CurrentValue()
		logger.LoggerInterface.Println("Setting Selection:", option, "as:", value)

		optionSelector := fmt.Sprintf(option, value)

		logger.LoggerInterface.Println(optionSelector)

		element, err2 := i.browser.Wd.FindElement("xpath", optionSelector)
		generalhelpers.Check(err2)

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

		return ElementCount(s, selector) <= 0
	}
}

func allExist(spec *spec, selectors []string) bool {
	for _, selector := range selectors {
		elements, err := spec.browser.Wd.FindElements("xpath", selector)
		generalhelpers.Check(err)

		if len(elements) <= 0 {
			return false
		}
	}

	return true
}

func KeepCheckingForElements(selectors []string, interval int, times int) func(*spec) bool {
	return func(spec *spec) bool {
		for i := 0; i < times; i++ {
			generalhelpers.Wait(interval)
			if allExist(spec, selectors) {
				return true
			}
		}

		return false
	}
}
