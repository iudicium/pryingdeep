package test_helpers

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/iudicium/pryingdeep/models"
)

func GetCrypto(db *gorm.DB, webPageID int) ([]models.Crypto, error) {
	var crypto []models.Crypto
	result := db.Where("web_page_id = ?", webPageID).Find(&crypto)
	if result.Error != nil {
		return nil, result.Error
	}

	return crypto, nil
}

func DeleteCryptoByWebPageId(db *gorm.DB, webPageId int) {
	db.Where("web_page_id = ?", webPageId).Delete(&models.Crypto{})
}

func GetPhoneNumbers(db *gorm.DB, webPageID int) ([]models.PhoneNumber, error) {
	var phoneNumbers []models.PhoneNumber
	result := db.Where("web_page_id = ?", webPageID).Find(&phoneNumbers)
	if result.Error != nil {
		return nil, result.Error
	}

	return phoneNumbers, nil
}

func DeletePhoneNumbersByCountryCode(db *gorm.DB, countryCode string) {
	db.Where("country_code = ?", countryCode).Delete(&models.PhoneNumber{})
}

func PreloadWebPage(db *gorm.DB, webPageID int) (*models.WebPage, error) {
	var webPageData models.WebPage

	if err := db.Preload(clause.Associations).Where("ID = ?", webPageID).Find(&webPageData).Error; err != nil {
		return nil, err
	}

	return &webPageData, nil

}

func CreateTestWebPage() error {
	db := models.GetDB()

	sql := "INSERT INTO web_pages (id, url, title, status_code, body, headers) VALUES (1, 'https://example.com', 'Example Page', 200, '<html><body>Hello, world!</body></html>', '{\"Content-Type\": \"text/html\"}')"

	result := db.Exec(sql)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
