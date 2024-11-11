package scanners

import (
	"fmt"
	"net/http"
)

type SnapchatScanner struct {
	username string
}

func (s *SnapchatScanner) Exists() bool {
	resp, err := http.Get(s.GetURL())
	if err != nil {
		return false
	}

	return resp.StatusCode != 404
}

func (s *SnapchatScanner) GetURL() string {
	return fmt.Sprintf("https://snapchat.com/add/%s", s.username)
}
