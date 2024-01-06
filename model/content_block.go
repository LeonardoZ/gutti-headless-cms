package model

import "time"

type ContentBlock struct {
	Id             int `gorm:"type:int;primarykey"`
	ExtId          string
	Title          string
	RawContent     string
	RawContentType string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (ContentBlock) TableName() string {
	return "content_blocks"
}
