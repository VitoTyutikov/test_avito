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

func (s *BannerService) IsBannerWithFeatureAndTagExists(featureID uint64, tagIDs []uint64, bannerID uint64) (bool, error) {
	return s.bannerRepo.IsBannerWithFeatureAndTagExistsWithoutBannerID(featureID, tagIDs, bannerID)
}

func (s *BannerService) Create(bannerReq *models.BannerRequestBody) (models.Banner, error) {
	return s.bannerRepo.Create(bannerReq)
}

func (s *BannerService) FindByID(id uint64) (*models.Banner, error) {
	return s.bannerRepo.FindByID(id)
}

func (s *BannerService) UpdateBanner(oldBanner *models.Banner, newBanner *models.BannerRequestBody) *gorm.DB {
	return s.bannerRepo.Update(oldBanner, newBanner)
}

func (s *BannerService) DeleteByID(id uint64) *gorm.DB {
	return s.bannerRepo.DeleteByID(id)
}
func (s *BannerService) GetBanners(featureID uint64, tagID uint64, limit int, offset int) ([]models.BannerResponseBody, error) {
	var banners []models.Banner
	query := s.bannerRepo.DB.Model(&models.Banner{})

	if featureID != 0 {
		query = query.Where("feature_id = ?", featureID)
	}
	if tagID != 0 {
		query = query.
			Joins("JOIN banner_tags ON banners.banner_id = banner_tags.banner_id").
			Where("banner_tags.tag_id = ?", tagID)
	}

	if limit != 0 {
		query = query.Limit(limit)
	}
	query = query.Offset(offset)

	if err := query.Find(&banners).Error; err != nil {
		return nil, err
	}

	bannerIDs := make([]uint64, len(banners))
	for i, banner := range banners {
		bannerIDs[i] = banner.BannerID
	}

	var bannerTags []models.BannerTag
	if err := s.bannerRepo.DB.Where("banner_id IN ?", bannerIDs).Find(&bannerTags).Error; err != nil {
		return nil, err
	}

	tagMap := make(map[uint64][]uint64)
	for _, tag := range bannerTags {
		tagMap[tag.BannerID] = append(tagMap[tag.BannerID], tag.TagID)
	}

	response := make([]models.BannerResponseBody, 0, len(banners))
	for _, banner := range banners {
		response = append(response, models.BannerResponseBody{
			BannerID:  banner.BannerID,
			TagIds:    tagMap[banner.BannerID],
			FeatureID: banner.FeatureID,
			Content:   banner.Content,
			IsActive:  banner.IsActive,
			CreatedAt: banner.CreatedAt,
			UpdatedAt: banner.UpdatedAt,
		})
	}

	return response, nil
}

//func (s *BannerService) GetBanners(featureID uint64, tagID uint64, limit int, offset int) ([]models.BannerResponseBody, error) {
//	var banners []models.Banner
//	query := s.bannerRepo.DB.Model(&models.Banner{})
//
//	if featureID != 0 {
//		query = query.Where("feature_id = ?", featureID)
//	}
//	if tagID != 0 {
//		query = query.
//			Joins("JOIN banner_tags ON banners.banner_id = banner_tags.banner_id").
//			Where("banner_tags.tag_id = ?", tagID)
//	}
//
//	if limit != 0 {
//		query = query.Limit(limit)
//	}
//	query = query.Offset(offset)
//
//	if err := query.Find(&banners).Error; err != nil {
//		return nil, err
//	}
//
//	var response []models.BannerResponseBody
//	for _, b := range banners {
//		var bannerTags []models.BannerTag
//		if err := s.bannerRepo.DB.Where("banner_id = ?", b.BannerID).
//			Find(&bannerTags).Error; err != nil {
//			return nil, err
//		}
//
//		tagIDs := make([]uint64, len(bannerTags))
//		for i, bt := range bannerTags {
//			tagIDs[i] = bt.TagID
//		}
//
//		response = append(response, models.BannerResponseBody{
//			BannerID:  b.BannerID,
//			TagIds:    tagIDs,
//			FeatureID: b.FeatureID,
//			Content:   b.Content,
//			IsActive:  b.IsActive,
//			CreatedAt: b.CreatedAt,
//			UpdatedAt: b.UpdatedAt,
//		})
//	}
//
//	return response, nil
//}

func (s *BannerService) GetUserBanner(featureId uint64, tagId uint64, token string) (models.Banner, error) {
	query := s.bannerRepo.DB.Model(&models.BannerTag{})
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

func (s *BannerService) CreateBannerWithTags(request *models.BannerRequestBody) (models.Banner, error) {
	return s.bannerRepo.CreateBannerWithTags(request)
}
