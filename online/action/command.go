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
		lg.Debug("Waiting:", wait, "milliseconds")
		generalhelpers.Wait(wait)
	}
}

func ExtractData(name string, selector string, attributeName string, channel chan *datastructures.Chunk) func(*interaction) {
	return func(i *interaction) {
		lg.Debug("Extracting contents from:", selector)
		channel <- &datastructures.Chunk{
			Name:  name,
			Value: i.browser.FindElementAttributes("xpath", selector, attributeName),
		}
	}
}

func ExtractStructuredData(name string, selector string, subSelectors map[string]*datastructures.Selector, channel chan *datastructures.Chunk) func(*interaction) {
	return func(i *interaction) {
		lg.Debug("Extracting structured contents from:", selector)

		containerElements, err := i.browser.Wd.FindElements("xpath", selector)
		generalhelpers.Check(err)
		data := make([]map[string]string, len(containerElements))

		for index, container := range containerElements {
			data[index] = make(map[string]string)
			for key, subSelector := range subSelectors {
				subElement, err := container.FindElement("xpath", subSelector.Selector)
				generalhelpers.Check(err)
				data[index][key], err = subElement.GetAttribute(subSelector.Attribute)
				generalhelpers.Check(err)
			}
		}

		channel <- &datastructures.Chunk{
			Name:  name,
			Value: data,
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
