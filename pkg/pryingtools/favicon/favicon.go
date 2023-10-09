package pryingtools

import (
	"fmt"
	"github.com/r00tk3y/prying-deep/pkg/utils"
	"io"
	"regexp"

	"github.com/twmb/murmur3"
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
				extractedUrls = append(extractedUrls, absoluteURL)
			} else {
				extractedUrls = append(extractedUrls, faviconURL)
			}
		}
	}
	return extractedUrls
}

func createMMH3Hash(data []byte) uint32 {
	hash := murmur3.New32()
	hash.Write(data)
	return hash.Sum32()
}

func fetchFaviconContent(url string, torProxy string) ([]byte, error) {
	//TODO: works for now but definetly need a better solution later on in life whenever i get a job
	torResult, _ := utils.CheckIfTorConnectionExists(torProxy)
	client := torResult.Client
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	faviconContent, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return faviconContent, nil
}

func GetFaviconHash(urls []string, torProxy string) []uint32 {
	var hashes []uint32
	for _, url := range urls {

		faviconContent, err := fetchFaviconContent(url, torProxy)
		if err != nil {
			fmt.Printf("Error fetching favicon from %s: %v\n", url, err)
			continue
		}

		hash := createMMH3Hash(faviconContent)
		hashes = append(hashes, hash)
	}
	fmt.Println(hashes)
	return hashes
}
