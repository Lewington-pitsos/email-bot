package emailbot

import (
	"time"
	"email-bot/offline/profile"
)

type Manager struct {
	bots []*Bot
	dataProfile *profile.DataProfile
}

func (m *Manager) AddBot(port int) {
	m.bots = append(m.bots, NewBot(port, m.dataProfile))
}

func (m *Manager) ScrapeAll() {
	for _, bot := range m.bots {
		go bot.Scrape()
	}

	time.Sleep(time.Millisecond * 10000)
}

func (m *Manager) DataProfile() {
	return m.profile
}

func NewManager() *Manager {
	return &Manager{
		bots: make([]*Bot, 0, 100),
		profile: profile.NewDataProfile()
	}
}
