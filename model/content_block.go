package model

import "time"

type ContentBlock struct {
	Id             int       `gorm:"type:int;primarykey;column:id"`
	ExtId          string    `gorm:"column:ext_id"`
	Title          string    `gorm:"column:title"`
	RawContent     string    `gorm:"column:raw_content"`
	RawContentType string    `gorm:"column:raw_content_type"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (ContentBlock) TableName() string {
	return "content_blocks"
}
