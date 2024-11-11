package scanners

import (
	"fmt"
	"time"

	"go_holmes/utils"
	"go_holmes/vars"

	"github.com/tebeka/selenium"
)

type TwitchScanner struct {
	username string
}

func (s *TwitchScanner) Exists() bool {
	// Get webdriver and service
	service, driver := utils.GetWebdriver(vars.ChromeDriverPath, vars.Headless)
	defer (*service).Stop()
	defer driver.Quit()

	// Load page
	driver.Get(s.GetURL())
	time.Sleep(vars.PAGE_LOAD_DELAY * time.Millisecond)

	// Check if the page shows a not found message
	_, err := driver.FindElement(selenium.ByXPATH, "//p[contains(text(), 'Sorry. Unless you')]")
	return err != nil
}

func (s *TwitchScanner) GetURL() string {
	return fmt.Sprintf("https://twitch.tv/%s", s.username)
}
