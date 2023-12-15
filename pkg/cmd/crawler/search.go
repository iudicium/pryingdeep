package crawler

import (
	"github.com/spf13/cobra"
)

var (
	keywords      []string
	searchEngines = []string{
		//Torch, not so much
		//"http://xmh57jrknzkhv6y3ls3ubitzfqnkrwxhopf5aygthi7d6rplyvk3noyd.onion/cgi-bin/omega/omega?P=",
		//Tordex works perfectly
		"http://tordexyb63aknnvuzyqabeqx6l7zdiesfos22nisv6zbj6c6o3h6ijyd.onion/search?query=",
		//	IceBerg. Not tested.
		//	http://iceberget6r64etudtzkyh5nanpdsqnkgav5fh72xtvry3jyu5u2r5qd.onion/
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
