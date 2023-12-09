package wordpress

import (
	"regexp"
	"strings"
)

// WordpressFinder is a type for finding related WordPress patterns in the given HTML.
type WordpressFinder struct {
	// Words is a list of words that identify WordPress patterns. These could be plugins, urls, etc.
	Words []string
	// Regex is the combined Words but with a regex pattern for searching
	Regex *regexp.Regexp
}

// NewWordpressPatternFinder creates a new WordpressPatternFinder instance with default configuration.
func NewWordpressPatternFinder() *WordpressFinder {
	words := []string{
		"wordpress",
		"wp-content",
		"wp-content/plugins",
		"wp-content/uploads",
		"wp-content/theme",
		"wp-includes",
	}

	pattern := `((?:<[^>]*>)?.*?)\b(` + strings.Join(words, "|") + `)\b((?:</[^>]*>)?.*?)`

	regex := regexp.MustCompile(pattern)

	return &WordpressFinder{
		Words: words,
		Regex: regex,
	}
}

// Find returns a list of matched HTML that contains the WordPress identifiers.
func (wpf *WordpressFinder) Find(html string) []string {
	matches := wpf.Regex.FindAllString(html, -1)
	return matches
}
