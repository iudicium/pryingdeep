package models

import "gorm.io/gorm"

type Response struct {
	gorm.Model
	URL        string      `json:"url" gorm:"uniqueIndex;not null"`
	StatusCode int         `json:"statusCode" gorm:"Index;not null"`
	Body       string      `json:"body" gorm:"type:text"`
	Context    PropertyMap `json:"context" gorm:"type:jsonb"`
	Headers    PropertyMap `json:"headers" gorm:"type:jsonb"`
}

func InsertResponse(statusCode int, body string, context PropertyMap, headers PropertyMap, url string) (uint, error) {
	response := &Response{
		StatusCode: statusCode,
		Body:       body,
		Context:    context,
		Headers:    headers,
		URL:        url,
	}

	result := db.Create(response)

	if result.Error != nil {
		return 0, result.Error
	}

	return response.ID, nil
}
