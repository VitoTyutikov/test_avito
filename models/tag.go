package models

type Tag struct {
	TagID       uint64 `json:"tag_id" gorm:"primarykey" `
	Description string `json:"description"`
}

type TagRequestBody struct {
	Description string `json:"description"`
}
