package crawler

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/iudicium/pryingdeep/configs"
)

var (
	keywords      []string
	searchEngines = []string{
		//Tordex works perfectly
		"http://tordexyb63aknnvuzyqabeqx6l7zdiesfos22nisv6zbj6c6o3h6ijyd.onion/search?query=",
		"http://orealmvxooetglfeguv2vp65a3rig2baq2ljc7jxxs4hsqsrcemkxcad.onion/search?query=",
	}
)

var SearchCMD = &cobra.Command{
	Use:   "search",
	Short: "Search different dark web search engines",
	Long:  "Search different dark web search engines using keywords or sentences to find the most accurate result.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg = setupCrawlerConfig(cmd)
		handleCrawlerTypeOptions(&cfg.Crawler, cmd)
		Crawl()

	},
}

func init() {
	SearchCMD.Flags().StringSliceVarP(&keywords, "keywords", "k", nil, "List of keywords or sentences for search")
	SearchCMD.Flags().MarkHidden("urls")

	SearchCMD.Flags().VisitAll(cli.ConfigureViper("crawler"))

}

// generateSearchURLS loops through the searchEngines variable and appends each keyword/sentence specified.
// which then changes the entrypoint of StartingURLS so that the "crawler" cfg startingURLS are overridden
func generateSearchURLS(keywords []string) {
	searchURLS := make([]string, 0)
	for _, keyword := range keywords {
		for _, engineURL := range searchEngines {
			url := engineURL + keyword
			searchURLS = append(searchURLS, url)

		}
	}
	cfg.Crawler.StartingURLS = searchURLS
}

// Is this necessary? Maybe for future flags in the search command
func handleCrawlerTypeOptions(c *configs.Crawler, cmd *cobra.Command) {
	if cmd.Flags().Changed("keywords") {
		c.Keywords = keywords
		generateSearchURLS(keywords)
	} else if len(c.Keywords) == 0 {
		fmt.Println(color.RedString("No keywords were provided while using the search command."))
		os.Exit(1)
	} else {
		generateSearchURLS(c.Keywords)
	}
}
