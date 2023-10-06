package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/configs"
)


var db *gorm.DB

type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time      `gorm:"autoCreateTime"`
  UpdatedAt time.Time      `gorm:"autoUpdateTime"` 
  DeletedAt gorm.DeletedAt `gorm:"index"`

}


func SetupDatabase(cfg *configs.DBConfig) {
	var err error
	db, err = gorm.Open(postgres.Open(cfg.DbURL), &gorm.Config{})

	if err != nil {
		logger.Fatal("models.Setup err: %v", err)
	}

	db.AutoMigrate(&Request{})
}