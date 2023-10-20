package parsers

import (
	"github.com/gocolly/colly/v2"
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/utils"
)

// TODO: Rename this file to something better, idk what yet, and also you probably do not need this as a package
// TODO: Instead you can just insert it into the crawler package || However keep in mind, the more
// TODO: modules you add the more you will have to expand it.
func ParseResponse(url string, body string, response *colly.Response) (int, error) {
	title, _ := utils.ExtractTitleFromBody(body)
	headers := utils.CreateMapFromValues(*response.Headers)

	ResId, err := models.CreatePage(
		url,
		title,
		response.StatusCode,
		body,
		headers,
	)
	if err != nil {
		return 0, err
	}

	return int(ResId), nil
}
