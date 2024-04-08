package repository

import (
	"avito_test_task/db"
	"avito_test_task/models"
	"gorm.io/gorm"
	"time"
)

type BannerRepository struct {
	DB *gorm.DB
}

func NewBannerRepository() *BannerRepository {
	return &BannerRepository{
		DB: db.DB,
	}
}
func (r *BannerRepository) FindAll() ([]models.Banner, error) {
	var banners []models.Banner
	err := r.DB.Find(&banners).Error
	return banners, err
}

func (r *BannerRepository) FindByID(bannerID uint) (*models.Banner, error) {
	var banner models.Banner
	err := r.DB.Where("banner_id = ?", bannerID).First(&banner).Error
	return &banner, err
}

func (r *BannerRepository) Update(oldBanner *models.Banner, newBanner *models.BannerRequestBody) error {
	return r.DB.Model(oldBanner).Updates(newBanner).Error
}

func (r *BannerRepository) Delete(bannerID uint) error {
	return r.DB.Delete(&models.Banner{}, bannerID).Error
}

func (r *BannerRepository) Create(bannerRequest *models.BannerRequestBody) error {
	banner := models.Banner{
		FeatureID: bannerRequest.FeatureID,
		Content:   bannerRequest.Content,
		IsActive:  bannerRequest.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return r.DB.Create(banner).Error
}
