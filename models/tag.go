package models

type Tag struct {
	TagID       uint   `json:"tag_id" gorm:"primarykey" `
	Description string `json:"description"`
}

type TagRequestBody struct {
	Description string `json:"description"`
}
