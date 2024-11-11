package scanners

import (
	"fmt"
	"time"

	"go_holmes/utils"
	"go_holmes/vars"

	"github.com/tebeka/selenium"
)

type TiktokScanner struct {
	username string
}

func (s *TiktokScanner) Exists() bool {
	// Get webdriver and service
	service, driver := utils.GetWebdriver(vars.ChromeDriverPath, vars.Headless)
	defer (*service).Stop()
	defer driver.Quit()

	// Load page
	driver.Get(s.GetURL())
	time.Sleep(vars.PAGE_LOAD_DELAY * time.Millisecond)

	// Check if the page shows a not found message
	_, err := driver.FindElement(selenium.ByXPATH, "//p[contains(text(), 'Couldn')]")
	return err != nil
}

func (s *TiktokScanner) GetURL() string {
	return fmt.Sprintf("https://tiktok.com/@%s", s.username)
}
