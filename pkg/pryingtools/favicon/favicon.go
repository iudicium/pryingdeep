package favicon

import (
	"fmt"
	"io"
	"regexp"

	"github.com/iudicium/pryingdeep/pkg/utils"
)

func ExtractFaviconUrls(html string, baseUrl string) []string {
	var extractedUrls []string
	faviconRegex := regexp.MustCompile(`<link[^>]*\srel=["'](icon|shortcut icon)["'][^>]*\shref=["']([^"']+)["']`)
	UrlValidationRegex := regexp.MustCompile(`^(?:[a-z+]+:)?//`)

	matches := faviconRegex.FindAllStringSubmatch(html, -1)
	for _, match := range matches {
		if len(match) >= 3 {
			faviconURL := match[2]
			if !UrlValidationRegex.MatchString(faviconURL) {
				absoluteURL := baseUrl + "/" + faviconURL
				fmt.Println(absoluteURL)
				extractedUrls = append(extractedUrls, absoluteURL)
			} else {
				extractedUrls = append(extractedUrls, faviconURL)
			}
		}
	}
	return extractedUrls
}

func fetchFaviconContent(url string, torProxy string) (string, error) {
	torResult, err := utils.CheckIfTorConnectionExists(torProxy)
	if err != nil {
		return "", err
	}

	response, err := torResult.Client.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	faviconContent, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	mm3hash := IconHash(faviconContent)
	return mm3hash, nil

}

func GetFaviconHash(urls []string, torProxy string) []string {
	var hashes []string
	for _, url := range urls {

		mm3, err := fetchFaviconContent(url, torProxy)
		if err != nil {

			fmt.Printf("Error fetching favicon from %s: %v\n", url, err)
			continue
		}

		hashes = append(hashes, mm3)
	}
	return hashes
}
