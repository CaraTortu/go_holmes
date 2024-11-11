package utils

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func randRange(min, max int) int {
	return min + rand.IntN(max-min)
}

func GetWebdriver(driverPath string, headless bool) (*selenium.Service, selenium.WebDriver) {
	// Set up ChromeDriver options
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// Configure Chrome-specific capabilities
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
			"--debug",
		},
	}

	if headless {
		chromeCaps.Args = append(chromeCaps.Args, "--headless=new")
	}

	caps.AddChrome(chromeCaps)

	// Start ChromeDriver service
	port := randRange(49152, 65535)
	service, err := selenium.NewChromeDriverService(driverPath, port, opts...)
	if err != nil {
		log.Fatal("Error starting ChromeDriver service:", err)
	}

	// Connect to ChromeDriver instance
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatal("Error creating WebDriver:", err)
	}

	return service, driver
}
