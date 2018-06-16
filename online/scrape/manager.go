package scrape

import (
	"github.com/tebeka/selenium"
)

type Manager struct {
	scrape  *Scrape
	browser selenium.WebDriver
}

func (m *Manager) AddAction() {

}

func NewManager(browser selenium.WebDriver) *Manager {
	return &Manager{
		browser: browser,
		scrape:  NewScrape(),
	}
}
