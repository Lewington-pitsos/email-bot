package	emailbot

type BotManagerInterface interface {
	AddBot(int)
	AddBots(int, int)
	Scrape()
	DataProfile()
	AddAction()
}