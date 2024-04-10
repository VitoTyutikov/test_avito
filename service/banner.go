package service

import (
	"avito_test_task/models"
	"avito_test_task/repository"
	"gorm.io/gorm"
)

type BannerService struct {
	bannerRepo *repository.BannerRepository
}

func NewBannerService() *BannerService {
	return &BannerService{
		bannerRepo: repository.NewBannerRepository(),
	}
}

func (s *BannerService) Create(bannerReq *models.BannerRequestBody) (models.Banner, error) {
	// Business logic can be added here before saving the banner
	return s.bannerRepo.Create(bannerReq)
}

func (s *BannerService) FindByID(id uint64) (*models.Banner, error) {
	// Additional business logic can be added here
	return s.bannerRepo.FindByID(id)
}

func (s *BannerService) UpdateBanner(oldBanner *models.Banner, newBanner *models.BannerRequestBody) *gorm.DB {
	return s.bannerRepo.Update(oldBanner, newBanner)
}

func (s *BannerService) DeleteByID(id uint64) *gorm.DB {
	// Additional logic before deletion, if necessary
	return s.bannerRepo.DeleteByID(id)
}

func (s *BannerService) FindAll() ([]models.Banner, error) {
	// Logic to handle retrieval of all banners
	return s.bannerRepo.FindAll()
}

func (s *BannerService) GetBanners(featureID uint64, tagID uint64, limit int, offset int) ([]models.BannerResponseBody, error) {
	var banners []models.Banner
	query := s.bannerRepo.DB.Model(&models.Banner{})

	if featureID != 0 {
		query = query.Where("feature_id = ?", featureID)
	}
	if tagID != 0 {
		query = query.Joins("JOIN banner_tags ON banners.banner_id = banner_tags.banner_id").Where("banner_tags.tag_id = ?", tagID)
	}

	if limit != 0 {
		query = query.Limit(limit)
	}
	query = query.Offset(offset)

	if err := query.Find(&banners).Error; err != nil {
		return nil, err
	}

	var response []models.BannerResponseBody
	for _, b := range banners {
		var bannerTags []models.BannerTag
		if err := s.bannerRepo.DB.Where("banner_id = ?", b.BannerID).Find(&bannerTags).Error; err != nil {
			return nil, err
		}

		tagIDs := make([]uint64, len(bannerTags))
		for i, bt := range bannerTags {
			tagIDs[i] = bt.TagID
		}

		response = append(response, models.BannerResponseBody{
			BannerID:  b.BannerID,
			TagIds:    tagIDs,
			FeatureID: b.FeatureID,
			Content:   b.Content,
			IsActive:  b.IsActive,
			CreatedAt: b.CreatedAt,
			UpdatedAt: b.UpdatedAt,
		})
	}

	return response, nil
}

func (s *BannerService) GetUserBanner(featureId uint64, tagId uint64, token string) (models.Banner, error) {
	query := s.bannerRepo.DB.Model(&models.BannerTag{}).Where("tag_id = ?", tagId)
	query = query.Joins("JOIN banners ON banner_tags.banner_id = banners.banner_id").Where("banners.feature_id = ?", featureId)
	if token == "user_token" {
		query = query.Where("banners.is_active = ?", true)
	}
	banner := models.Banner{}
	if err := query.First(&banner).Error; err != nil {
		return banner, err
	}
	return banner, nil

}
