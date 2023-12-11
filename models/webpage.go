package models

// PageData is embedded into the webPage model for structured data
type PageData struct {
	URL        string      `json:"url" gorm:"uniqueIndex;not null"`
	Title      string      `json:"title" gorm:"Index"`
	StatusCode int         `json:"statusCode" gorm:"Index;not null"`
	Body       string      `json:"body" gorm:"type:text"`
	Headers    PropertyMap `json:"headers" gorm:"type:jsonb"`
}
type WebPage struct {
	Model
	PageData `json:"pageData"`
	//References are here for exporting structured data from the webPage model.
	Emails       *Email              `json:"email" gorm:"foreignKey:WebPageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PhoneNumbers *PhoneNumber        `json:"phoneNumbers" gorm:"foreignKey:WebPageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Crypto       *Crypto             `json:"crypto" gorm:"foreignKey:WebPageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Wordpress    *WordpressFootPrint `json:"wordpress" gorm:"foreignKey:WebPageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func CreatePage(url string, title string, statusCode int, body string, headers PropertyMap) (uint, error) {
	pageData := PageData{
		URL:        url,
		Title:      title,
		StatusCode: statusCode,
		Body:       body,
		Headers:    headers,
	}

	webPage := WebPage{
		PageData: pageData,
	}
	// BUG: ERROR: invalid byte sequence for encoding "UTF8": 0xfc (SQLSTATE 22021)
	if err := db.Create(&webPage).Error; err != nil {
		return 0, err
	}

	return webPage.ID, nil
}
