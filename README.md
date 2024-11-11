# Go holmes
A username scanning tool to see if accounts with the same username appear on other social media.

## Requirements
- Chromedriver & Chrome. You can download chromedriver from [here](https://developer.chrome.com/docs/chromedriver/downloads)

## How to run
1. ```sh
    $ make build
    ```

2. ```sh
    $ ./go_holmes -h
    Web scraping tool to find social media accounts

    Usage:
      go_holmes [flags]

    Flags:
      -c, --chrome-driver-path string   Path to the ChromeDriver executable (default "binaries/chromedriver")
          --headless                    Run Chrome in headless mode
      -h, --help                        help for go_holmes
      -u, --username string             Username for the social media account to search for 
```


