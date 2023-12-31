package models

import (
	"gorm.io/gorm"

	"github.com/iudicium/pryingdeep/pkg/logger"
)

type PhoneNumber struct {
	Model

	//WebPageID is the serves as a foreign key to web_pages
	WebPageID           int    `json:"pageId"`
	InternationalNumber string `json:"internationalNumber" gorm:"uniqueIndex"`
	NationalNumber      string `json:"nationalNumber" gorm:"unique"`
	CountryCode         string `json:"countryCode"`
}

func CreatePhoneNumber(webPageID int, interNum, natNum string, code string) error {
	phoneNumber := &PhoneNumber{
		WebPageID:           webPageID,
		InternationalNumber: interNum,
		NationalNumber:      natNum,
		CountryCode:         code,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(phoneNumber).Error; err != nil {
			logger.Errorf("Error during transaction: %s", err)
			tx.Rollback()
			return err
		}
		logger.Infof("Created phone number record | Number: %s, CountryCode: %v", interNum, code)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
