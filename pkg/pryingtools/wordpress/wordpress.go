package wordpress

import (
	"fmt"
	"regexp"
)

func FindWordpressPatterns(html string, url string) ([]string, error) {
	words := []string{
		"wp-content",
		"wp-content/plugins",
		"wp-content/uploads",
		"wp-content/theme",
		"wp-includes",
	}

	//TODO: what if a website contains external wordpress links? That could be actually useful as it will
	//TODO: allow for analysis of links

	pattern := "\\b(" + regexp.QuoteMeta(words[0])
	for i := 1; i < len(words); i++ {
		pattern += "|" + regexp.QuoteMeta(words[i])
	}
	pattern += ")\\b"
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil, err
	}

	matches := regex.FindAllString(html, -1)
	return matches, nil
}
