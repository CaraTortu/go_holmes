package scraper

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"go_holmes/scraper/scanners"
	"go_holmes/vars"

	"github.com/mitchellh/colorstring"
	"github.com/schollz/progressbar/v3"
	"github.com/zenthangplus/goccm"
)

const STATUS_MESSAGE = "[blue][i] [white]Found [red]%d [white]accounts..."

func GetAccountsWithUsername() {
	accountsFound := []string{}
	accountsFoundMutex := sync.Mutex{}
	c := goccm.New(vars.MAX_WEB_INSTANCES)

	// Get scanners
	scannerList := scanners.GetScanners(vars.Username)
	linksScannerList := scanners.GetLinksScanners(vars.Username)

	// Get progress bar
	bar := progressbar.Default(int64(len(scannerList)+len(linksScannerList)), colorstring.Color(fmt.Sprintf(STATUS_MESSAGE, 0)))

	// Start timer
	start := time.Now()

	// Get accounts scanners
	for _, s := range scannerList {
		c.Wait()

		go func(s scanners.AccountScanner) {
			defer func() {
				c.Done()
				bar.Add(1)
				bar.Describe(colorstring.Color(fmt.Sprintf(STATUS_MESSAGE, len(accountsFound))))
			}()

			if s.Exists() {
				// Save the account found
				accountsFoundMutex.Lock()
				defer accountsFoundMutex.Unlock()

				accountsFound = append(accountsFound, s.GetURL())
			}
		}(s)
	}

	// Get account links scanners
	for _, s := range linksScannerList {
		c.Wait()

		go func(s scanners.AccountLinksScanner) {
			defer func() {
				c.Done()
				bar.Add(1)
				bar.Describe(colorstring.Color(fmt.Sprintf(STATUS_MESSAGE, len(accountsFound))))
			}()

			links := s.GetLinks()

			for _, link := range links {
				accountsFound = append(accountsFound, link)
			}
		}(s)
	}
	// Wait for all scanners to finish
	c.WaitAllDone()
	bar.Finish()

	if len(accountsFound) == 0 {
		colorstring.Printf("[red][-] [white]No accounts found for %s. Took %d seconds\n", vars.Username, time.Since(start)/time.Second)
		return
	}

	// Filter out duplicates
	uniqueAccountsFound := make([]string, 0, len(accountsFound))
	for _, account := range accountsFound {
		found := false
		for _, uniqueAccount := range uniqueAccountsFound {
			if account == uniqueAccount {
				found = true
				break
			}
		}

		if !found {
			uniqueAccountsFound = append(uniqueAccountsFound, account)
		}
	}

	// Sort by alphabetical order
	sort.Strings(uniqueAccountsFound)

	// Print accounts found
	colorstring.Println(fmt.Sprintf("[green][+] [white]%d Accounts found in %d seconds: ", len(uniqueAccountsFound), time.Since(start)/time.Second))
	for _, account := range uniqueAccountsFound {
		fmt.Println("  -", account)
	}
}
