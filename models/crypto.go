package models

import (
	"fmt"

	"github.com/lib/pq"
)

// Crypto stores various cryptographic findings, such as
type Crypto struct {
	Model
	//WebPageID is the serves as a foreign key to web_pages
	WebPageID    int            `json:"pageId"`
	PGPKeys      pq.StringArray `json:"PGPKey" gorm:"type:text[]"`
	Certificates pq.StringArray `json:"Certificates" gorm:"type:text[]"`
	Wallets      pq.StringArray `json:"Wallets" gorm:"type:text[]"`
}

func CryptoCreate(c Crypto) (Crypto, error) {
	result := db.Create(&c)
	if result.Error != nil {
		fmt.Println(result.Error)
		return c, result.Error
	}
	return c, nil
}
