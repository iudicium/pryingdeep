package parsers

import (
	"fmt"
	"go.uber.org/zap"

	"github.com/gocolly/colly/v2"

	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/pkg/utils"
)

func ParseRequest(request *colly.Request) (uint, error) {
	logger.Info("Parsing request data...")
	url := request.URL.String()
	headers := utils.CreateMapFromValues(*request.Headers)
	fmt.Println(headers)

	ctx := utils.ConvertContextToPropertyMap(*request.Ctx)

	reqId, err := models.InsertRequest(
		url,
		headers,
		ctx,
		request.Depth,
		request.Method,
		request.ResponseCharacterEncoding,
		request.ProxyURL,
	)
	if err != nil {
		logger.Error("Error saving request to database.", zap.Error(err))
	}
	return reqId, nil

}

// https://stackoverflow.com/questions/59244393/how-to-handle-nullable-postgres-jsonb-data-and-parse-it-as-json
