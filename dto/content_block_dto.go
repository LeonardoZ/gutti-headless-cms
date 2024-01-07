package dto

import "time"

type UpsertContentBlockDTO struct {
	Title          string `binding:"required,max=500,min=1" json:"title"`
	RawContent     string `binding:"required,min=1" json:"raw_content"`
	RawContentType string `binding:"required,max=500,min=1" json:"raw_content_type"`
}

type ContentBlockDTO struct {
	ExtId          string
	Title          string
	RawContent     string
	RawContentType string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
