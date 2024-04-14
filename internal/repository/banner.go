package repository

import (
	"avito_test_task/internal/db"
	"avito_test_task/internal/models"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
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

func (r *BannerRepository) FindByID(bannerID uint64) (*models.Banner, error) {
	var banner models.Banner
	err := r.DB.Where("banner_id = ?", bannerID).Find(&banner).Error
	return &banner, err
}

func (r *BannerRepository) DeleteByID(bannerID uint64) *gorm.DB {
	return r.DB.Delete(&models.Banner{}, bannerID)
}

func (r *BannerRepository) CreateBannerWithTags(request *models.BannerRequestBody) (models.Banner, error) {
	tx := r.DB.Session(&gorm.Session{SkipDefaultTransaction: true}).Begin()

	banner := models.Banner{
		FeatureID: request.FeatureID,
		Content:   request.Content,
		IsActive:  *request.IsActive,
	}
	if err := tx.Create(&banner).Error; err != nil {
		tx.Rollback()
		return banner, err
	}

	bannerTags := make([]models.BannerTag, len(request.TagIds))
	for i, tagID := range request.TagIds {
		bannerTags[i] = models.BannerTag{
			BannerID: banner.BannerID,
			TagID:    tagID,
		}
	}

	if err := tx.Create(&bannerTags).Error; err != nil {
		tx.Rollback()
		return banner, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return banner, err
	}

	return banner, nil
}

func (r *BannerRepository) UpdateBannerWithTags(oldBanner *models.Banner, newBanner *models.BannerRequestBody) (int, error) {
	tx := r.DB.Session(&gorm.Session{SkipDefaultTransaction: true}).Begin()

	oldBanner.FeatureID = newBanner.FeatureID
	oldBanner.Content = newBanner.Content
	oldBanner.IsActive = *newBanner.IsActive
	oldBanner.UpdatedAt = time.Now()
	if err := tx.Save(oldBanner).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	resultDelete := tx.Where("banner_id = ?", oldBanner.BannerID).Delete(&models.BannerTag{})
	if resultDelete.Error != nil {
		tx.Rollback()
		return http.StatusInternalServerError, resultDelete.Error
	}
	if resultDelete.RowsAffected == 0 {
		tx.Rollback()
		return http.StatusNotFound, nil
	}

	bannerTags := make([]models.BannerTag, len(newBanner.TagIds))
	for i, tagID := range newBanner.TagIds {
		bannerTags[i] = models.BannerTag{
			BannerID: oldBanner.BannerID,
			TagID:    tagID,
		}
	}
	if err := tx.Create(&bannerTags).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	if err := tx.Commit().Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (r *BannerRepository) GetUserBanner(featureId uint64, tagId uint64, token string) (models.Banner, error) {
	query := r.DB.Model(&models.BannerTag{})
	query = query.
		Select("banners.*").
		Joins("JOIN banners ON banner_tags.banner_id = banners.banner_id").
		Where("banner_tags.tag_id = ?", tagId).
		Where("banners.feature_id = ?", featureId)

	if token == "user_token" {
		query = query.Where("banners.is_active = ?", true)
	}
	var banner models.Banner
	query = query.Find(&banner)
	if query.Error != nil {
		return banner, query.Error
	}
	if query.RowsAffected == 0 {
		return banner, gorm.ErrRecordNotFound
	}
	return banner, nil
}

func (r *BannerRepository) GetBanners(featureID uint64, tagID uint64, limit int, offset int) ([]models.BannerResponseBody, error) {
	var results []struct {
		models.Banner
		TagIDs string
	}
	query := r.DB.Model(&models.Banner{}).
		Select("banners.*, json_agg(banner_tags.tag_id::bigint) as tag_ids").
		Joins("join banner_tags on banners.banner_id = banner_tags.banner_id").
		Where("? = 0 OR banners.feature_id = ?", featureID, featureID).
		Where("? = 0 OR banner_tags.tag_id = ?", tagID, tagID).
		Group("banners.banner_id")

	query = query.Offset(offset)
	if limit != 0 {
		query = query.Limit(limit)
	}

	if err := query.Scan(&results).Error; err != nil {
		return nil, err
	}
	response := make([]models.BannerResponseBody, len(results))
	for i, result := range results {
		var tagIDs []uint64
		if err := json.Unmarshal([]byte(result.TagIDs), &tagIDs); err != nil {
			return nil, err
		}
		response[i] = models.BannerResponseBody{
			BannerID:  result.BannerID,
			TagIds:    tagIDs,
			FeatureID: result.FeatureID,
			Content:   result.Content,
			IsActive:  result.IsActive,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
	}

	return response, nil
}
