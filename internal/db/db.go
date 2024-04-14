package db

import (
	"avito_test_task/internal/models"
	"encoding/json"
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
	insertTestData()
	return nil
}

func insertTestData() {
	var count int64
	DB.Model(&models.Tag{}).Count(&count)
	if count > 0 {
		return
	}
	DB.Model(&models.Feature{}).Count(&count)
	if count > 0 {
		return
	}
	for i := 0; i < 1000; i++ {
		if err := DB.Create(&models.Tag{}).Error; err != nil {
			panic(err)
		}
		if err := DB.Create(&models.Feature{}).Error; err != nil {
			panic(err)
		}
	}

	DB.Model(models.Banner{}).Count(&count)
	if count > 0 {
		return
	}

	for i := 0; i < 1000; i++ {
		isActive := !(i%10 == 0)
		_, err := CreateBannerWithTags(&models.BannerRequestBody{
			TagIds:    []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			FeatureID: uint64(i) + 1,
			Content: json.RawMessage(`{
    "title": "some_title",
    "text": "some_text",
    "url": "some_url"
  }`),
			IsActive: &isActive,
		})
		if err != nil {
			panic(err)
			return
		}
	}

}

func CreateBannerWithTags(request *models.BannerRequestBody) (models.Banner, error) {
	tx := DB.Session(&gorm.Session{SkipDefaultTransaction: true}).Begin()
	//tx := DB.Begin()
	banner := models.Banner{
		FeatureID: request.FeatureID,
		Content:   request.Content,
		IsActive:  *request.IsActive,
	}
	if err := tx.Create(&banner).Error; err != nil {
		tx.Rollback()
		return banner, err
	}
	for _, tagID := range request.TagIds {
		bannerTag := models.BannerTag{
			BannerID: banner.BannerID,
			TagID:    tagID,
		}
		if err := tx.Create(&bannerTag).Error; err != nil {
			tx.Rollback()
			return banner, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return banner, err
	}
	return banner, nil
}
