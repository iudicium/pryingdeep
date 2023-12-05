package wordpress

import (
	"fmt"
	"regexp"
	"strings"
)

// TODO ad post processing and struct methods
func FindWordpressPatterns(html string) ([]string, error) {
	words := []string{
		"wordpress",
		"wp-content",
		"wp-content/plugins",
		"wp-content/uploads",
		"wp-content/theme",
		"wp-includes",
	}

	pattern := `((?:<[^>]*>)?.*?)\b(` + strings.Join(words, "|") + `)\b((?:</[^>]*>)?.*?)`

	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil, err
	}

	matches := regex.FindAllString(html, -1)

	return matches, nil
}
