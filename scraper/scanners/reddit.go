package scanners

import (
	"fmt"
	"time"

	"go_holmes/utils"
	"go_holmes/vars"

	"github.com/tebeka/selenium"
)

type RedditScanner struct {
	username string
}

func (s *RedditScanner) Exists() bool {
	// Get webdriver and service
	service, driver := utils.GetWebdriver(vars.ChromeDriverPath, vars.Headless)
	defer (*service).Stop()
	defer driver.Quit()

	// Load page
	driver.Get(s.GetURL())
	time.Sleep(vars.PAGE_LOAD_DELAY * time.Millisecond)

	// Check if the page shows a not found message
	_, err := driver.FindElement(selenium.ByXPATH, "//span[contains(text(), 'Sorry,')]")
	return err != nil
}

func (s *RedditScanner) GetURL() string {
	return fmt.Sprintf("https://reddit.com/u/%s", s.username)
}
