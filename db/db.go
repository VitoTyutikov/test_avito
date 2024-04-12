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
		//SkipDefaultTransaction: true,
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

	//for i := 0; i < 500; i++ {
	//	if err := DB.Create(&models.Tag{Description: "tag_" + strconv.Itoa(i+1)}).Error; err != nil {
	//		panic(err)
	//	}
	//	if err := DB.Create(&models.Feature{Description: "feature_" + strconv.Itoa(i+1)}).Error; err != nil {
	//		panic(err)
	//	}
	//}
	//for i := 500; i < 1000; i++ {
	//	if err := DB.Create(&models.Tag{Description: "tag_" + strconv.Itoa(i+1)}).Error; err != nil {
	//		panic(err)
	//	}
	//	if err := DB.Create(&models.Feature{Description: "feature_" + strconv.Itoa(i+1)}).Error; err != nil {
	//		panic(err)
	//	}
	//}

	return nil
}
