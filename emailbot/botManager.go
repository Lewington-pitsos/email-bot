package emailbot

import (
	"email-bot/offline/profile"
	"time"
)

type Manager struct {
	bots        []*Bot
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

func (m *Manager) DataProfile() *profile.DataProfile {
	return m.dataProfile
}

func NewManager() *Manager {
	return &Manager{
		bots:        make([]*Bot, 0, 100),
		dataProfile: profile.NewDataProfile(),
	}
}
