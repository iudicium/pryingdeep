package models

import "gorm.io/gorm/clause"

type WebPage struct {
	Model
	URL          string              `json:"url" gorm:"uniqueIndex;not null"`
	Title        string              `json:"title" gorm:"Index"`
	StatusCode   int                 `json:"statusCode" gorm:"Index;not null"`
	Body         string              `json:"body" gorm:"type:text"`
	Headers      PropertyMap         `json:"headers" gorm:"type:jsonb"`
	Emails       *Email              `json:"email" gorm:"foreignKey:WebPageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PhoneNumbers *PhoneNumber        `json:"phoneNumbers" gorm:"foreignKey:WebPageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Crypto       *Crypto             `json:"crypto" gorm:"foreignKey:WebPageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Wordpress    *WordpressFootPrint `json:"wordpress" gorm:"foreignKey:WebPageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

func PreloadWebPage(webPageID int) (*WebPage, error) {
	var webPageData WebPage

	if err := db.Preload(clause.Associations).Where("ID = ?", webPageID).Find(&webPageData).Error; err != nil {
		return nil, err
	}

	return &webPageData, nil

}
