package models

type Tag struct {
	TagID uint64 `json:"tag_id" gorm:"primarykey" `
}
