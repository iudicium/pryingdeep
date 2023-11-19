package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	pryingtools "github.com/iudicium/pryingdeep/pkg/pryingtools/favicon"
)

var filePath string
var fileContents []byte

func init() {
	filePath = filepath.Join("data", "favicon.html")

	var err error
	fileContents, err = os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
}

func TestFaviconExtraction(t *testing.T) {
	HtmlString := string(fileContents)
	baseUrl := "http://explorerzydxu5ecjrkwceayqybizmpjjznk5izmitf2modhcusuqlid.onion"
	torProxy := "socks5://localhost:9050"
	urls := pryingtools.ExtractFaviconUrls(HtmlString, baseUrl)
	if len(urls) != 1 {
		t.Errorf("there's only 1 favicon in that baseUrl")
	}

	hashes := pryingtools.GetFaviconHash(urls, torProxy)
	fmt.Println(hashes)

}
