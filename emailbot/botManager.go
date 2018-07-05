package emailbot

import (
	"email-bot/offline/profile"
	"email-bot/online/action"
	"email-bot/online/scrape"
	"time"
)

type Manager struct {
	bots          []*Bot
	profile   *profile.ProfileInterface
	instructions  []action.Action
	scrapeProfile *scrape.Profile
}

func (m *Manager) AddBot(port int) {
	m.bots = append(m.bots, NewBot(port, m.profile, m.scrapeProfile.Instructions()))
}

func (m *Manager) ScrapeAll() {
	for _, bot := range m.bots {
		go bot.Scrape()
	}

	time.Sleep(time.Millisecond * 90000)
}

func (m *Manager) AddAction(critical bool) *action.Action {
	action := action.NewAction(critical)
	m.scrapeProfile.AddAction(action)
	return action
}

func (m *Manager)AddProfile(profileInstance profile.ProfileInterface) *Manager {
	m.profile = profileInstance
	return m
}

func NewManager() *Manager {
	return &Manager{
		bots:          make([]*Bot, 0, 100),
		scrapeProfile: scrape.NewProfile(),
	}
}
