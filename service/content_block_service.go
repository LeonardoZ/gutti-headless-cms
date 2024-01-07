package service

import (
	"github.com/LeonardoZ/gutti-headless-cms/dto"
)

type ContentBlockService interface {
	Create(dto dto.UpsertContentBlockDTO)
	Update(id int, dto dto.UpsertContentBlockDTO)
	Delete(id int)
	FindById(id int) dto.ContentBlockDTO
	FindByExtId(uuid string) dto.ContentBlockDTO
	FindAll() []dto.ContentBlockDTO
}
