package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/fatih/color"

	"github.com/pryingbytez/pryingdeep/models"
)

const checkTor string = "https://check.torproject.org/api/ip"

type TorCheckResult struct {
	IsTor  bool
	Client *http.Client // Include the HTTP client
}

// SetupNewTorCLient  allows us to pass custom transport, used for checking
// If our tor connection actually works.
// TODO: change it to colly.Request
func SetupNewTorClient(torProxy string) (*http.Client, error) {
	torProxyUrl, err := url.Parse(torProxy)

	if err != nil {
		return nil, err
	}

	//TODO: add support for passing custom timeouts
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 15}

	return client, nil
}

// TODO: Rename this function to something better because it also returns a client with a tor connection

func CheckIfTorConnectionExists(torProxy string) (*TorCheckResult, error) {
	client, err := SetupNewTorClient(torProxy)
	if err != nil {
		return nil, err
	}

	resp, _ := client.Get(checkTor)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data struct {
		IsTor bool `json:"IsTor"`
		IP    string
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	fmt.Println(color.GreenString("[+]"), "Connected to Tor IP:", color.RedString(data.IP))
	result := &TorCheckResult{
		IsTor:  data.IsTor,
		Client: client,
	}

	return result, nil
}

// CreateMapFromValues is used to convert a map into a PropertyMap which is
// JSONB in postgres
func CreateMapFromValues(data map[string][]string) models.PropertyMap {
	resultMap := make(models.PropertyMap)

	for key, values := range data {
		if len(values) == 1 {
			resultMap[key] = values[0]
		} else {
			resultMap[key] = values
		}
	}

	return resultMap
}

// CompileRegex - compiles various Collly filters such as URLFIlters, DomainFilters, etc.
func CompileRegex(patterns []string) ([]*regexp.Regexp, error) {
	regexSlice := make([]*regexp.Regexp, len(patterns))
	for i, pattern := range patterns {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		regexSlice[i] = regex
	}
	return regexSlice, nil
}

// ExtractTitleFromBody is there because I could not parse it with Colly so
// There's a nice regexp, it works.
func ExtractTitleFromBody(body string) (string, error) {
	titleRegex := regexp.MustCompile(`(?i)<title[^>]*>([^<]+)</title>`)
	matches := titleRegex.FindStringSubmatch(body)

	if len(matches) >= 2 {
		return matches[1], nil
	}

	return "", fmt.Errorf("no title found in the HTML body")
}
