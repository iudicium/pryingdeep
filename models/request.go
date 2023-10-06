package models

import (
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model

	URL        string      `json:"url" gorm:"Index;not null"`
	Headers    PropertyMap `json:"headers" gorm:"type:jsonb"`
	Context    PropertyMap `json:"context" gorm:"type:jsonb"`
	Depth      int         `json:"depth"`
	Method     string      `json:"method" gorm:"type:varchar(30)"`
	Body       []byte      `json:"body" gorm:"type:bytea"`
	ChEncoding string      `json:"chEncoding" gorm:"type:varchar(255)"`
	ProxyURL   string      `json:"proxyurl" gorm:"type:varchar(255)"`
}

func InsertRequest(url string, headers PropertyMap, ctx PropertyMap, depth int, method string, chEncoding string, proxyURL string) (uint, error) {
	req := &Request{
		URL:        url,
		Headers:    headers,
		Context:    ctx,
		Depth:      depth,
		Method:     method,
		ChEncoding: chEncoding,
		ProxyURL:   proxyURL,
	}

	result := db.Create(req)

	if result.Error != nil {
		return 0, result.Error
	}

	return req.ID, nil
}
