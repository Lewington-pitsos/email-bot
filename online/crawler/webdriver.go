package crawler

import (
	"fmt"

	"github.com/tebeka/selenium"
)

func NewWebDriver(portNum int) selenium.WebDriver {
	const (
		// These paths will be different on your system.
		seleniumPath    = "/usr/local/bin/selenium-server-standalone-3.12.0.jar"
		geckoDriverPath = "/usr/local/bin/geckodriver"
	)

	selenium.SetDebug(true)
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", portNum))
	if err != nil {
		panic(err)
	}

	return wd
}
