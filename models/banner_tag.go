package models

// BannerTag TODO: maybe add feature to bannerTag to do search faster
type BannerTag struct {
	BannerID uint64 `gorm:"primaryKey;autoIncrement:false"`
	TagID    uint64 `gorm:"primaryKey;autoIncrement:false"`
	Banner   Banner `gorm:"references:BannerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tag      Tag    `gorm:"references:TagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
