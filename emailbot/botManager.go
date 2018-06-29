package emailbot

import (
	"email-bot/offline/profile"
	"email-bot/online/action"
	"email-bot/online/scrape"
	"time"
)

type Manager struct {
	bots          []*Bot
	dataProfile   *profile.DataProfile
	instructions  []action.Action
	scrapeProfile *scrape.Profile
}

func (m *Manager) AddBot(port int) {
	m.bots = append(m.bots, NewBot(port, m.dataProfile, m.scrapeProfile.Instructions()))
}

func (m *Manager) AddBots(startingPort int, number int) {
	for i := 0; i <= number; i++ {
		m.AddBot(startingPort + i)
	}
}

func (m *Manager) Scrape() {
	for _, bot := range m.bots {
		go bot.Scrape()
	}

	time.Sleep(time.Millisecond * 10000)
}

func (m *Manager) AddAction(critical bool) *action.Action {
	action := action.NewAction(critical)
	m.scrapeProfile.AddAction(action)
	return action
}

func (m *Manager) DataProfile() *profile.DataProfile {
	return m.dataProfile
}

func NewManager() *Manager {
	return &Manager{
		bots:          make([]*Bot, 0, 100),
		dataProfile:   profile.NewDataProfile(),
		scrapeProfile: scrape.NewProfile(),
	}
}
