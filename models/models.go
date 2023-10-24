package models

import (
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	loger "gorm.io/gorm/logger"
)

var db *gorm.DB

type Model struct {
	ID uint `gorm:"primaryKey"`
}

func SetupDatabase(dbUrl string) *gorm.DB {
	var err error
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: loger.Default.LogMode(loger.Info),
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
