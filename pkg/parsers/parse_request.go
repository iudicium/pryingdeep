package parsers

import (
	"context"
	"fmt"
	"encoding/json"
	"go.uber.org/zap"
	"github.com/gocolly/colly/v2"
    "github.com/lib/pq"

	"github.com/r00tk3y/prying-deep/internal/database/sql"
)

func ParseRequest(request *colly.Request, logger *zap.Logger)  {

	context := context.Background()

	url := request.URL.String()
    fmt.Println(request.Headers)
	headers, err := json.Marshal(*request.Headers)
	if err != nil {
		logger.Error("Problem during extraction of headers from request:", zap.Error(err))
	}
	colly_context, err :=  json.Marshal(*request.Ctx)	
	if err != nil {
		logger.Error("Problem found during extraction of context from request:", zap.Error(err))
	}
    depth := request.Depth
    method := request.Method
    // body := request.Body
    responseCharacterEncoding := request.ResponseCharacterEncoding
    proxyURL := request.ProxyURL
    logger.Info("Request Data",
        zap.String("URL", url),
        zap.Any("Headers", headers),
        // zap.Any("Ctx", ctx),
        zap.Int("Depth", depth),
        zap.String("Method", method),
        // zap.ByteString("Body", body),
        zap.String("ResponseCharacterEncoding", responseCharacterEncoding),
        zap.String("ProxyURL", proxyURL),
    )

	db_request_id := sql.CreateRequest(context,
		sql.CreateRequestParams{
			Url: url,
			Headers: headers,
			Ctx: pq.NullRawMessage{Valid: true, RawMessage: []byte(ctx)},

		},
	)
}