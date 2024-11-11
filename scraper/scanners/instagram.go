package scanners

import (
	"fmt"
	"time"

	"go_holmes/utils"
	"go_holmes/vars"

	"github.com/tebeka/selenium"
)

type InstagramScanner struct {
	username string
}

func (s *InstagramScanner) Exists() bool {
	// Get webdriver and service
	service, driver := utils.GetWebdriver(vars.ChromeDriverPath, vars.Headless)
	defer (*service).Stop()
	defer driver.Quit()

	// Load page
	driver.Get(s.GetURL())
	time.Sleep(vars.PAGE_LOAD_DELAY * time.Millisecond)

	// Check if we are being rate limited
	_, err := driver.FindElement(selenium.ByXPATH, "//span[contains(text(), 'Something went wrong')]")
	if err == nil {
		return false
	}

	// Check if allow cookies button is present
	allowCookiesButton, err := driver.FindElement(selenium.ByXPATH, "//button[contains(text(), 'Allow all cookies')]")
	if err == nil {
		allowCookiesButton.Click()
		time.Sleep(2500 * time.Millisecond)
	}

	// Check if the page shows a not found message
	_, err = driver.FindElement(selenium.ByXPATH, "//span[contains(text(), 'Sorry')]")
	return err != nil
}

func (s *InstagramScanner) GetURL() string {
	return fmt.Sprintf("https://www.instagram.com/%s", s.username)
}
