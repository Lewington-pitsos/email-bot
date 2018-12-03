package browser

import (
	"email-bot/helpers/generalhelpers"

	"github.com/tebeka/selenium"
)

type Browser struct {
	Wd      selenium.WebDriver
	service *selenium.Service
}

func (b *Browser) Quit() {
	b.service.Stop()
	b.Wd.Quit()
}

func (b *Browser) FindElementAttributes(selectMode string, selector string, attributeName string) []string {
	elements, err := b.Wd.FindElements(selectMode, selector)
	generalhelpers.Check(err)
	attributes := make([]string, len(elements))

	for index, element := range elements {
		attribute, err := element.GetAttribute(attributeName)
		generalhelpers.Check(err)
		attributes[index] = attribute
	}

	return attributes
}

func NewBrowser(portNum int) *Browser {
	browser, service := NewWebDriver(portNum)

	return &Browser{
		Wd:      browser,
		service: service,
	}
}
