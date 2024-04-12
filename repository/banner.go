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

func (r *BannerRepository) IsBannerWithFeatureAndTagExistsWithoutBannerID(featureID uint64, tagIDs []uint64, bannerID uint64) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Banner{}).
		Joins("INNER JOIN banner_tags ON banners.banner_id = banner_tags.banner_id").
		Where("banners.feature_id = ?", featureID).
		Where("banner_tags.tag_id IN (?)", tagIDs).
		Where("banners.banner_id != ?", bannerID).
		Where("banner_tags.banner_id != ?", bannerID).
		Count(&count).Error
	return count > 0, err
}

func (r *BannerRepository) FindAll() ([]models.Banner, error) {
	var banners []models.Banner
	err := r.DB.Find(&banners).Error
	return banners, err
}

func (r *BannerRepository) FindByID(bannerID uint64) (*models.Banner, error) {
	var banner models.Banner
	err := r.DB.Where("banner_id = ?", bannerID).First(&banner).Error
	return &banner, err
}

func (r *BannerRepository) Update(oldBanner *models.Banner, newBanner *models.BannerRequestBody) *gorm.DB {
	//return r.DB.Model(oldBanner).Updates(newBanner).Error
	oldBanner.FeatureID = newBanner.FeatureID
	oldBanner.Content = newBanner.Content
	oldBanner.IsActive = *newBanner.IsActive
	oldBanner.UpdatedAt = time.Now()
	return r.DB.Save(&oldBanner)
}

func (r *BannerRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

func (r *BannerRepository) Commit() *gorm.DB {
	return r.DB.Commit()
}

func (r *BannerRepository) DeleteByID(bannerID uint64) *gorm.DB {
	return r.DB.Delete(&models.Banner{}, bannerID)
}

func (r *BannerRepository) Create(bannerRequest *models.BannerRequestBody) (models.Banner, error) {
	banner := models.Banner{
		FeatureID: bannerRequest.FeatureID,
		Content:   bannerRequest.Content,
		IsActive:  *bannerRequest.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := r.DB.Create(&banner).Error
	return banner, err
}

func (r *BannerRepository) CreateBannerWithTags(request *models.BannerRequestBody) (models.Banner, error) {
	tx := r.DB.Begin()
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

//func (r*BannerRepository) UpdateBannerWithTags(oldBanner *models.Banner, newBanner *models.BannerRequestBody) (int, error) {
//	tx:=r.DB.Begin()
//	resultUpdate := b.bannerService.UpdateBanner(oldBanner, &request)
//	if resultUpdate.Error != nil {
//		//c.JSON(http.StatusInternalServerError, gin.H{"error": resultUpdate.Error.Error()})
//		return http.StatusInternalServerError, nil
//	}
//
//	resultDelete := b.bannerTagService.DeleteByBannerID(oldBanner.BannerID)
//
//	if resultDelete.Error != nil {
//		//c.JSON(http.StatusInternalServerError, gin.H{"error": resultDelete.Error.Error()})
//		return http.StatusInternalServerError, nil
//	}
//	if resultDelete.RowsAffected == 0 {
//		//c.Status(http.StatusNotFound)
//		return http.StatusNotFound, nil
//	}
//	bannerExists, err := b.bannerService.IsBannerWithFeatureAndTagExists(request.FeatureID, request.TagIds)
//	if bannerExists {
//		//c.JSON(http.StatusBadRequest, gin.H{"error": "banner with this tag_id and feature_id already exists"})
//		return http.StatusBadRequest, nil
//	}
//	if err != nil {
//		//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return http.StatusInternalServerError, err
//	}
//
//	for _, tagID := range newBanner.TagIds {
//		if err = b.bannerTagService.Create(&models.BannerTag{
//			BannerID: oldBanner.BannerID,
//			TagID:    tagID,
//		}); err != nil {
//			//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return http.StatusInternalServerError,nil
//		}
//	}
//	if err = tx.Commit().Error; err != nil {
//		tx.Rollback()
//		//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return http.StatusInternalServerError,nil
//	}
//	return http.StatusOK, nil
//
//}
