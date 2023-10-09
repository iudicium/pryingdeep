package models

type WebPage struct {
	Model
	URL        string      `json:"url" gorm:"uniqueIndex;not null"`
	Title      string      `json:"title" gorm:"Index"`
	StatusCode int         `json:"statusCode" gorm:"Index;not null"`
	Body       string      `json:"body" gorm:"type:text"`
	Headers    PropertyMap `json:"headers" gorm:"type:jsonb"`
}

func CreatePage(url string, title string, statusCode int, body string, headers PropertyMap) (uint, error) {
	webpage := &WebPage{
		URL:        url,
		Title:      title,
		StatusCode: statusCode,
		Body:       body,
		Headers:    headers,
	}

	result := db.Create(webpage)

	if result.Error != nil {
		return 0, result.Error
	}

	return webpage.ID, nil
}
