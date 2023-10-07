package parsers

import (
	"github.com/gocolly/colly/v2"
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/utils"
)

func ParseResponse(response *colly.Response) (uint, error) {
	url := response.Request.URL.String()

	// Body might take a long time but we thugging it out
	body := string(response.Body)
	context := utils.ConvertContextToPropertyMap(*response.Ctx)
	headers := utils.CreateMapFromValues(*response.Headers)

	ResId, err := models.InsertResponse(
		response.StatusCode,
		body,
		context,
		headers,
		url,
	)
	if err != nil {
		return 0, err
	}

	return ResId, nil
}
