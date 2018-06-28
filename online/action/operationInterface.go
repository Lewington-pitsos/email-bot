package action

type operationInterface interface {
	clone() *operationInterface
	AddCommand()
	addBrowser(browser *browser.Browser)
}