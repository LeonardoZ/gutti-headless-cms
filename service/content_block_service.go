package service

import (
	"github.com/LeonardoZ/gutti-headless-cms/dto"
)

type ContentBlockService interface {
	Create(dto dto.UpsertContentBlockDTO) error
	Update(id int, dto dto.UpsertContentBlockDTO) error
	Delete(id int) error
	FindById(id int) (*dto.ContentBlockDTO, error)
	FindByExtId(uuid string) (*dto.ContentBlockDTO, error)
	FindAll() ([]dto.ContentBlockDTO, error)
}
