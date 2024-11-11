package scanners

type AccountScanner interface {
	Exists() bool
	GetURL() string
}

type AccountLinksScanner interface {
	GetLinks() []string
	GetURL() string
}

func GetScanners(username string) []AccountScanner {
	return []AccountScanner{
		&InstagramScanner{username: username},
		&XScanner{username: username},
		&TiktokScanner{username: username},
		&RedditScanner{username: username},
		&OnlyfansScanner{username: username},
		&FacebookScanner{username: username},
		&SnapchatScanner{username: username},
		&TwitchScanner{username: username},
		&PinterestScanner{username: username},
		&VcsoScanner{username: username},
	}
}

func GetLinksScanners(username string) []AccountLinksScanner {
	return []AccountLinksScanner{
		&LinktreeScanner{username: username},
	}
}
