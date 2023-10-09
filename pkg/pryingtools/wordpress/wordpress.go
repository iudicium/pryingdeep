package wordpress

import (
	"fmt"
	"regexp"
	"strings"
)

func FindWordpressPatterns(html string, url string) ([]string, error) {
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
	//TODO: there's too much html being returned which leads to dupliccates with WebPage model and ineffciency.
	//FIXME: need a fix
	matches := regex.FindAllString(html, -1)

	return matches, nil
}

//func findAllUrlsInHtml(html string) (string, error) {
//	urlPattern := `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
//
//	regex, err := regexp.Compile(urlPattern)
//	if err != nil {
//		fmt.Println("Error compiling regex:", err)
//		return "", err
//	}
//
//	matches := regex.FindAllString(html, -1)
//	result := strings.Join(matches, " ")
//	return result, nil
//}
