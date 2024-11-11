package scanners

import (
	"fmt"
	"time"

	"go_holmes/utils"
	"go_holmes/vars"

	"github.com/tebeka/selenium"
)

type LinktreeScanner struct {
	username string
}

func (s *LinktreeScanner) GetLinks() []string {
	// Get webdriver and service
	service, driver := utils.GetWebdriver(vars.ChromeDriverPath, vars.Headless)
	defer (*service).Stop()
	defer driver.Quit()

	hrefs := []string{}

	if err := driver.Get(s.GetURL()); err != nil {
		return hrefs
	}

	// Wait for the page to load properly
	time.Sleep(vars.PAGE_LOAD_DELAY * time.Millisecond)

	// There is an element with the data-test attribute set to "MessageCopy" in the page
	// that contains the text
	_, err := driver.FindElement(selenium.ByXPATH, "//h2[@data-testid='MessageCopy']")
	if err == nil {
		return hrefs
	}

	// Account exists, add it onto hrefs
	hrefs = append(hrefs, s.GetURL())

	// Find all links in the page
	links, err := driver.FindElements(selenium.ByXPATH, "//a[@data-testid='LinkButton']")
	if err != nil {
		return hrefs
	}

	// Get the href attribute of each link
	for _, link := range links {
		href, err := link.GetAttribute("href")
		if err != nil {
			continue
		}
		hrefs = append(hrefs, href)
	}

	return hrefs
}

func (s *LinktreeScanner) GetURL() string {
	return fmt.Sprintf("https://linktr.ee/%s", s.username)
}
