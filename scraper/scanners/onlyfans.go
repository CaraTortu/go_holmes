package scanners

import (
	"fmt"
	"time"

	"go_holmes/utils"
	"go_holmes/vars"

	"github.com/tebeka/selenium"
)

type OnlyfansScanner struct {
	username string
}

func (s *OnlyfansScanner) Exists() bool {
	// Get webdriver and service
	service, driver := utils.GetWebdriver(vars.ChromeDriverPath, vars.Headless)
	defer (*service).Stop()
	defer driver.Quit()

	// Load page
	driver.Get(s.GetURL())
	time.Sleep(vars.PAGE_LOAD_DELAY * time.Millisecond)

	// Check if the page shows a not found message
	_, err := driver.FindElement(selenium.ByXPATH, "//h1[contains(text(), 'Sorry')]")
	return err != nil
}

func (s *OnlyfansScanner) GetURL() string {
	return fmt.Sprintf("https://onlyfans.com/%s", s.username)
}
