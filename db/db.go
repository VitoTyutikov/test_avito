package db

import (
	"avito_test_task/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	dsn := "host=localhost user=postgres password=postgres dbname=test_avito port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&models.Tag{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&models.Feature{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&models.Banner{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&models.BannerTag{})
	if err != nil {
		return err
	}
	return nil
}
