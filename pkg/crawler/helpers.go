package crawler

import (
	"github.com/gocolly/colly/v2"

	"github.com/pryingbytez/pryingdeep/models"
	"github.com/pryingbytez/pryingdeep/pkg/utils"
)

// ParseResponse creates a record in the database for web_pages
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
