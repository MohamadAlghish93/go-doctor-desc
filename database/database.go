package database

import (
	"doc-desc/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.User{}, &models.Receipt{}, &models.Medicine{}, &models.ReceiptMedicine{}, &models.Information{}, &models.Role{})
}
