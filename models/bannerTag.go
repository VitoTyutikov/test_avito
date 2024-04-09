package models

//type BannerTag struct {
//	BannerID uint   `gorm:"primaryKey;autoIncrement:false"`
//	TagID    uint   `gorm:"primaryKey;autoIncrement:false"`
//	Banner   Banner `gorm:"foreignKey:BannerID;references:BannerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
//	Tag      Tag   `gorm:"foreignKey:TagID;references:TagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
//}

type BannerTag struct {
	BannerID uint   `gorm:"primaryKey;autoIncrement:false"`
	TagID    uint   `gorm:"primaryKey;autoIncrement:false"`
	Banner   Banner `gorm:"references:BannerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tag      Tag    `gorm:"references:TagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
