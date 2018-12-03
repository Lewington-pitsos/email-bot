package action

import (
	"email-bot/datastructures"
	"email-bot/helpers/generalhelpers"
	"email-bot/lg"
	"fmt"
)

// +---------------------------------------------------------------------------------------+
//										INTERACTIONS
// +---------------------------------------------------------------------------------------+

func VisitPage(url string) func(*interaction) {
	return func(i *interaction) {
		lg.Debug("Visiting:", url)
		i.browser.Wd.Get(url)
	}
}

func FillField(selector string) func(*interaction) {
	return func(i *interaction) {
		element, err := i.browser.Wd.FindElement("xpath", selector)
		generalhelpers.Check(err)

		value := i.CurrentValue()
		lg.Debug("Filling element:", selector, "with:", value)
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
		lg.Debug("Setting Selection:", option, "as:", value)

		optionSelector := fmt.Sprintf(option, value)

		lg.Debug(optionSelector)

		element, err2 := i.browser.Wd.FindElement("xpath", optionSelector)
		generalhelpers.Check(err2)

		element.Click()
	}
}

func Click(selector string) func(*interaction) {
	return func(i *interaction) {
		lg.Debug("Clicking element:", selector)
		element, err := i.browser.Wd.FindElement("xpath", selector)
		generalhelpers.Check(err)

		element.Click()
	}
}

func Wait(wait int) func(*interaction) {
	return func(i *interaction) {
		lg.Debug("Waiting:", wait, "miliseconds")
		generalhelpers.Wait(wait)
	}
}

func ExtractData(name, string, selector string, attributeName string, channel chan datastructures.Signal) func(*interaction) {
	return func(i *interaction) {
		lg.Debug("Extracting contents from:", selector)
		attributes := i.browser.FindElementAttributes("xpath", selector, attributeName)
		for _, attribute := range attributes {
			channel <- datastructures.Signal{
				Name:  name,
				Value: attribute,
			}
		}

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
		lg.Debug("Checking for element:", selector)

		return ElementCount(s, selector) > 0
	}
}

func CheckDoesntExist(selector string) func(*spec) bool {
	return func(s *spec) bool {
		lg.Debug("Confirming missing element:", selector)

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
