package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/pryingbytez/pryingdeep/pkg/logger"
)

var db *gorm.DB

type Model struct {
	ID uint `gorm:"primaryKey"`
}

func SetupDatabase(dbUrl string) *gorm.DB {
	var err error
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	db.Debug()

	if err != nil {
		logger.Fatalf("models.Setup err: %v", err)
	}

	err = db.AutoMigrate(&WebPage{}, &WordpressFootPrint{}, &Email{}, &PhoneNumber{}, &Crypto{})
	if err != nil {
		logger.Errorf("error during AutoMigrations", err)
	}
	return db
}
