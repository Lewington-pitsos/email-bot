package browser

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
)

func NewWebDriver(portNum int) (selenium.WebDriver, *selenium.Service) {
	const (
		seleniumPath    = "/usr/local/bin/selenium-server-standalone-3.12.0.jar"
		geckoDriverPath = "/usr/local/bin/chromedriver"
	)

	opts := []selenium.ServiceOption{
		selenium.GeckoDriver(geckoDriverPath),
		selenium.Output(os.Stderr),
	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, portNum, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}

	selenium.SetDebug(true)
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", portNum))
	if err != nil {
		panic(err)
	}

	return wd, service
}
