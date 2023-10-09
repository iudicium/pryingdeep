package tests

import (
	"github.com/r00tk3y/prying-deep/pkg/pryingtools/wordpress"
	"os"
	"path/filepath"
	"testing"
)

var filePath string
var fileContents []byte

func init() {
	filePath = filepath.Join("data", "wordpress.html")

	var err error
	fileContents, err = os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
}

func TestWordpressPatternsInHtml(t *testing.T) {
	url := "https://based.win/"
	HtmlString := string(fileContents)

	matches, err := wordpress.FindWordpressPatterns(HtmlString, url)
	if err != nil {
		t.Error("something went wrong during test of wordpress", err)
	}

	if len(matches) != 247 {
		t.Errorf("regexp erro")
	}
}

func TestNoWordpressPattern(t *testing.T) {
	url := "https://based.win/"
	filePath = filepath.Join("data", "no_wordpress.html")

	fileContents, _ = os.ReadFile(filePath)

	matches, _ := wordpress.FindWordpressPatterns(string(fileContents), url)
	if len(matches) != 0 {
		t.Errorf("expected array length should be 1")
	}

}
