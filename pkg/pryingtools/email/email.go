package email

import (
	"regexp"
)

func FindEmail(html string) []string {
	//TODO: add database logic, will need to see in which URL the email has been found
	emailRegex := `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`

	re := regexp.MustCompile(emailRegex)

	matches := re.FindAllString(html, -1)

	return matches
}
