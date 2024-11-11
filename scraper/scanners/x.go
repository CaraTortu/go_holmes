package scanners

import (
	"fmt"
	"time"

	"go_holmes/utils"
	"go_holmes/vars"

	"github.com/tebeka/selenium"
)

type XScanner struct {
	username string
}

func (s *XScanner) Exists() bool {
	// Get webdriver and service
	service, driver := utils.GetWebdriver(vars.ChromeDriverPath, vars.Headless)
	defer (*service).Stop()
	defer driver.Quit()

	// Load page
	driver.Get(s.GetURL())
	driver.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		// Check if the page contains the username
		_, err := wd.FindElement(selenium.ByXPATH, fmt.Sprintf("//span[contains(text(), '%s')]", s.username))
		return err == nil, nil
	}, time.Millisecond*vars.PAGE_LOAD_DELAY)

	// Check if the page shows a not found message
	_, err := driver.FindElement(selenium.ByXPATH, "//span[contains(text(), 'This account doesn')]")
	return err != nil
}

func (s *XScanner) GetURL() string {
	return fmt.Sprintf("https://x.com/%s", s.username)
}
