package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gormLog "gorm.io/gorm/logger"

	"github.com/iudicium/pryingdeep/pkg/logger"
)

var db *gorm.DB

type Model struct {
	ID uint `gorm:"primaryKey"`
}

// SetupDatabase also auto migrates the models if there are any changes
func SetupDatabase(dbUrl string) *gorm.DB {
	var err error
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		TranslateError: true,
		Logger:         gormLog.Default.LogMode(gormLog.Silent),
	})

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

func GetDB() *gorm.DB {
	return db
}
