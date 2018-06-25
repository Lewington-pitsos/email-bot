package emailbot

import (
	"email-bot/offline/profile"
	"email-bot/online/action"
	"time"
)

type Manager struct {
	bots        []*Bot
	dataProfile *profile.DataProfile
	instructions []action.Action
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

func (m *Manager)AddAction() *action.Action{
	action := action.NewAction()
	m.instructions := append(m.instructions, action)
	return action
}

func (m *Manager) DataProfile() *profile.DataProfile {
	return m.dataProfile
}

func NewManager() *Manager {
	return &Manager{
		bots:        make([]*Bot, 0, 100),
		dataProfile: profile.NewDataProfile(),
		instructions: male([]action.Action, 0, 40),
	}
}
