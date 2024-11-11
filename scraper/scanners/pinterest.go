package scanners

import (
	"fmt"
	"time"

	"go_holmes/utils"
	"go_holmes/vars"
)

type PinterestScanner struct {
	username string
}

func (s *PinterestScanner) Exists() bool {
	// Get webdriver and service
	service, driver := utils.GetWebdriver(vars.ChromeDriverPath, vars.Headless)
	defer (*service).Stop()
	defer driver.Quit()

	// Load page
	driver.Get(s.GetURL())
	time.Sleep(vars.PAGE_LOAD_DELAY * time.Millisecond)

	// If the URL is https://www.pinterest.com/?show_error=true it means the account does not exist
	url, err := driver.CurrentURL()
	return err == nil && url != "https://www.pinterest.com/?show_error=true"
}

func (s *PinterestScanner) GetURL() string {
	return fmt.Sprintf("https://pinterest.com/%s", s.username)
}
