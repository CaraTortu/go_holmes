package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go_holmes/scraper"
	"go_holmes/vars"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_holmes",
	Short: "Web scraping tool to find social media accounts",
	Run: func(cmd *cobra.Command, args []string) {
		// Search for social media account
		scraper.GetAccountsWithUsername()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&vars.Headless, "headless", false, "Run Chrome in headless mode")
	rootCmd.PersistentFlags().StringVarP(&vars.ChromeDriverPath, "chrome-driver-path", "c", "binaries/chromedriver", "Path to the ChromeDriver executable")
	rootCmd.Flags().StringVarP(&vars.Username, "username", "u", "", "Username for the social media account to search for")

	rootCmd.MarkFlagRequired("username")
}
