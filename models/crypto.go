package models

import (
	"fmt"
	"github.com/lib/pq"
)

type Crypto struct {
	Model
	WebPageId    int            `json:"pageId"`
	WebPage      WebPage        `json:"page" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PGPKeys      pq.StringArray `json:"PGPKey" gorm:"type:text[]"`
	Certificates pq.StringArray `json:"Certificates" gorm:"type:text[]"`
}

func CryptoCreate(c Crypto) (Crypto, error) {
	result := db.Create(&c)
	if result.Error != nil {
		fmt.Println(result.Error)
		return c, result.Error
	}
	return c, nil
}

func GetCrypto(webPageID int) ([]Crypto, error) {
	var crypto []Crypto
	result := db.Where("web_page_id = ?", webPageID).Find(&crypto)
	if result.Error != nil {
		return nil, result.Error
	}

	return crypto, nil
}

func DeleteCryptoByWebPageId(webPageId int) {
	db.Where("web_page_id = ?", webPageId).Delete(&Crypto{})
}
