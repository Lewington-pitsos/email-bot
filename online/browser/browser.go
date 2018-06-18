package browser

import "github.com/tebeka/selenium"

type Browser struct {
	Wd      selenium.WebDriver
	service *selenium.Service
}

func (b *Browser) Quit() {
	b.service.Stop()
	b.Wd.Quit()
}

func NewBrowser(portNum int) *Browser {
	browser, service := NewWebDriver(portNum)

	return &Browser{
		Wd:      browser,
		service: service,
	}
}
