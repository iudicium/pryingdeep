package models

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/pkg/logger"
)

var db *gorm.DB

type Model struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func SetupDatabase(cfg *configs.DBConfig) {
	var err error
	db, err = gorm.Open(postgres.Open(cfg.DbURL), &gorm.Config{})

	if err != nil {
		logger.Fatalf("models.Setup err: %v", err)
	}

	err = db.AutoMigrate(&Response{})
	if err != nil {
		logger.Errorf("error during AutoMigrations", err)
	}

}
