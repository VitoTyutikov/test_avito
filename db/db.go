package db

import (
	"avito_test_task/errors"
	"avito_test_task/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=test_avito port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errors.CheckError(err)
	DB.AutoMigrate(&models.Banner{})
	DB.AutoMigrate(&models.Tags{})
	DB.AutoMigrate(&models.Features{})
	DB.AutoMigrate(&models.BannerTag{})
}
