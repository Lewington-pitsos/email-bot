package emailbot

import (
	"email-bot/offline/profile"
	"email-bot/online/action"
	"email-bot/online/scrape"
)

type Manager struct {
	bots          []*Bot
	profile       profile.ProfileInterface
	instructions  []action.Action
	scrapeProfile *scrape.Profile
}

func (m *Manager) AddBot(port int) {
	botProfile := m.profile.Generate().Values()
	m.bots = append(m.bots, NewBot(port, botProfile, m.scrapeProfile.Instructions(), m.profile.Saveable()))
}

func (m *Manager) AddBots(port int, number int) {
	for i := 0; i < number; i++ {
		m.AddBot(port + i)
	}
}

func (m *Manager) ScrapeAll() {
	for _, bot := range m.bots {
		go bot.Scrape()
	}
}

func (m *Manager) AddAction(critical bool) *action.Action {
	action := action.NewAction(critical)
	m.scrapeProfile.AddAction(action)
	return action
}

func (m *Manager) SetProfile(profileInstance profile.ProfileInterface) *Manager {
	m.profile = profileInstance
	return m
}

func NewManager() *Manager {
	return &Manager{
		bots:          make([]*Bot, 0, 100),
		scrapeProfile: scrape.NewProfile(),
	}
}
