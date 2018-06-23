package emailbot

import (
	"time"
)

type Manager struct {
	bots []*Bot
}

func (m *Manager) AddBot(port int) {
	m.bots = append(m.bots, NewBot(port))
}

func (m *Manager) ScrapeAll() {
	for _, bot := range m.bots {
		go bot.Scrape()
	}

	time.Sleep(time.Millisecond * 10000)
}

func NewManager() *Manager {
	return &Manager{
		bots: make([]*Bot, 0, 100),
	}
}
