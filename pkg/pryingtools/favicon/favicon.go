package pryingtools

import (
	"encoding/base64"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/iudicium/pryingdeep/pkg/utils"

	"github.com/twmb/murmur3"
)

// Sadly, I do not have neough money for api-keys for shodan,so we will have to search for better ways of deanonymizaiton
func ExtractFaviconUrls(html string, baseUrl string) []string {
	var extractedUrls []string
	faviconRegex := regexp.MustCompile(`<link[^>]*\srel=["'](icon|shortcut icon)["'][^>]*\shref=["']([^"']+)["']`)
	UrlValidationRegex := regexp.MustCompile(`^(?:[a-z+]+:)?//`)

	matches := faviconRegex.FindAllStringSubmatch(html, -1)
	fmt.Println(matches)
	for _, match := range matches {
		if len(match) >= 3 {
			faviconURL := match[2]
			fmt.Println(faviconURL)
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

func createMMH3Hash(data string) uint32 {
	hash := murmur3.New32()
	hash.Write([]byte(data))
	fmt.Println(hash.Sum32())

	return hash.Sum32()
}

func fetchFaviconContent(url string, torProxy string) (string, error) {
	torResult, err := utils.CheckIfTorConnectionExists(torProxy)
	if err != nil {
		return "", err
	}

	client := torResult.Client
	fmt.Println(url)
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	faviconContent, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	base64Encoded := base64.StdEncoding.EncodeToString(faviconContent)
	base64Encoded = strings.TrimRight(base64Encoded, "=")

	return base64Encoded, nil
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
	return hashes
}
