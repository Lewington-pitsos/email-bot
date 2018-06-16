package online

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

func Test() {
	const (
		// These paths will be different on your system.
		seleniumPath    = "/usr/local/bin/selenium-server-standalone-3.12.0.jar"
		geckoDriverPath = "/usr/local/bin/geckodriver"
		port            = 8081
	)

	opts := []selenium.ServiceOption{
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("http://google.com"); err != nil {
		panic(err)
	}

	time.Sleep(time.Millisecond * 20000)
}
