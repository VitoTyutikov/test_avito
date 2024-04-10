package db

import (
	"avito_test_task/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() error {
	dsn := "host=localhost user=postgres password=postgres dbname=test_avito port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	if err = DB.AutoMigrate(&models.Tag{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.Feature{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.Banner{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.BannerTag{}); err != nil {
		return err
	}

	//for i := 0; i < 20; i++ {
	//	DB.Create(&models.Tag{Description: "tag_" + strconv.Itoa(i+1)})
	//	DB.Create(&models.Feature{Description: "feature_" + strconv.Itoa(i+1)})
	//}

	return nil
}
