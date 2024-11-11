package scanners

import (
	"fmt"
	"time"

	"go_holmes/utils"
	"go_holmes/vars"

	"github.com/tebeka/selenium"
)

type FacebookScanner struct {
	username string
}

func (s *FacebookScanner) Exists() bool {
	// Get webdriver and service
	service, driver := utils.GetWebdriver(vars.ChromeDriverPath, vars.Headless)
	defer (*service).Stop()
	defer driver.Quit()

	// Load page
	driver.Get(s.GetURL())
	time.Sleep(vars.PAGE_LOAD_DELAY * time.Millisecond)

	// Check if we need to click on cookies
	el, err := driver.FindElements(selenium.ByXPATH, "//span[contains(text(), 'Allow all cookies')]")
	if err == nil && len(el) > 1 {
		el[1].Click()
		time.Sleep(500 * time.Millisecond)
	}

	// Check if the page shows a not found message
	_, err = driver.FindElement(selenium.ByXPATH, "//span[contains(text(), 'This content isn')]")
	return err != nil
}

func (s *FacebookScanner) GetURL() string {
	return fmt.Sprintf("https://facebook.com/%s", s.username)
}
