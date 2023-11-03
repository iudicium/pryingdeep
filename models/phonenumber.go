package models

import (
	"github.com/pryingbytez/pryingdeep/pkg/logger"
	"gorm.io/gorm"
)

type PhoneNumber struct {
	Model
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

func GetPhoneNumbers(webPageID int) ([]PhoneNumber, error) {
	var phoneNumbers []PhoneNumber
	result := db.Where("web_page_id = ?", webPageID).Find(&phoneNumbers)
	if result.Error != nil {
		return nil, result.Error
	}

	return phoneNumbers, nil
}

func DeletePhoneNumbersByCountryCode(countryCode string) {
	db.Where("country_code = ?", countryCode).Delete(&PhoneNumber{})
}
