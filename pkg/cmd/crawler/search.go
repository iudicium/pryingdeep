package crawler

import (
	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		cfg = setupCrawlerConfig(cmd, "search")
		Crawl()

	},
}

func init() {
	initCrawler(SearchCMD, "search")

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
