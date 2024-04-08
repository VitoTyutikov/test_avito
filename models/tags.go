package models

type Tags struct {
	TagID       uint   `json:"tag_id" gorm:"primarykey"`
	Description string `json:"description"`
}

type TagsRequestBody struct {
	Description string `json:"description"`
}
