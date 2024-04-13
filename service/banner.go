package service

import (
	"avito_test_task/models"
	"avito_test_task/repository"
	"errors"
	"gorm.io/gorm"
	"net/http"
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

func (s *BannerService) FindByID(id uint64) (*models.Banner, error) {
	return s.bannerRepo.FindByID(id)
}

func (s *BannerService) UpdateBannerWithTags(oldBanner *models.Banner, newBanner *models.BannerRequestBody) (int, error) {
	return s.bannerRepo.UpdateBannerWithTags(oldBanner, newBanner)
}

func (s *BannerService) DeleteByID(id uint64) *gorm.DB {
	return s.bannerRepo.DeleteByID(id)
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
//	bannerIDs := make([]uint64, len(banners))
//	for i, banner := range banners {
//		bannerIDs[i] = banner.BannerID
//	}
//
//	var bannerTags []models.BannerTag
//	if err := s.bannerRepo.DB.Where("banner_id IN ?", bannerIDs).Find(&bannerTags).Error; err != nil {
//		return nil, err
//	}
//
//	tagMap := make(map[uint64][]uint64)
//	for _, tag := range bannerTags {
//		tagMap[tag.BannerID] = append(tagMap[tag.BannerID], tag.TagID)
//	}
//
//	response := make([]models.BannerResponseBody, 0, len(banners))
//	for _, banner := range banners {
//		response = append(response, models.BannerResponseBody{
//			BannerID:  banner.BannerID,
//			TagIds:    tagMap[banner.BannerID],
//			FeatureID: banner.FeatureID,
//			Content:   banner.Content,
//			IsActive:  banner.IsActive,
//			CreatedAt: banner.CreatedAt,
//			UpdatedAt: banner.UpdatedAt,
//		})
//	}
//
//	return response, nil
//}

func (s *BannerService) GetBanners(featureID uint64, tagID uint64, limit int, offset int) ([]models.BannerResponseBody, error) {
	return s.bannerRepo.GetBanners(featureID, tagID, limit, offset)
}

func (s *BannerService) GetUserBanner(featureId uint64, tagId uint64, token string) (models.Banner, error) {
	return s.bannerRepo.GetUserBanner(featureId, tagId, token)
}

func (s *BannerService) CreateBannerWithTags(request *models.BannerRequestBody) (models.Banner, error) {
	return s.bannerRepo.CreateBannerWithTags(request)
}

func (s *BannerService) Update(id uint64, request *models.BannerRequestBody) (int, error) {
	oldBanner, err := s.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, nil
	} else if err != nil {
		return http.StatusInternalServerError, err
	}
	// check no new tags which exists with feature_id added to banner
	bannerExists, err := s.IsBannerWithFeatureAndTagExists(request.FeatureID, request.TagIds, oldBanner.BannerID)
	if bannerExists {
		return http.StatusBadRequest, errors.New("banner with this tag_id and feature_id already exists")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return s.UpdateBannerWithTags(oldBanner, request)

}
