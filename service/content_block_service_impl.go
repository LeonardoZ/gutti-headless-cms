package service

import (
	"github.com/LeonardoZ/gutti-headless-cms/dto"
	"github.com/LeonardoZ/gutti-headless-cms/model"
	"github.com/LeonardoZ/gutti-headless-cms/repository"
	"github.com/go-playground/validator/v10"
)

type ContentBlockServiceImpl struct {
	ContentBlockRepository repository.ContentBlockRepository
	Validate               *validator.Validate
}

func NewContentBlockService(contentRepository repository.ContentBlockRepository, validate *validator.Validate) ContentBlockService {
	return &ContentBlockServiceImpl{
		ContentBlockRepository: contentRepository,
		Validate:               validate,
	}
}

// Create implements ContentBlockService.
func (s *ContentBlockServiceImpl) Create(dto dto.UpsertContentBlockDTO) {
	err := s.Validate.Struct(dto)
	if err != nil {
		panic(err)
	}
	model := model.ContentBlock{
		Title:          dto.Title,
		RawContent:     dto.RawContent,
		RawContentType: dto.RawContentType,
	}

	s.ContentBlockRepository.Save(model)
}

// Delete implements ContentBlockService.
func (s *ContentBlockServiceImpl) Delete(id int) {
	s.ContentBlockRepository.Delete(id)
}

// FindAll implements ContentBlockService.
func (s *ContentBlockServiceImpl) FindAll() []dto.ContentBlockDTO {
	result := s.ContentBlockRepository.FindAll()
	responses := []dto.ContentBlockDTO{}
	for _, value := range result {
		dto := dto.ContentBlockDTO{
			RawContent:     value.RawContent,
			RawContentType: value.RawContentType,
			Title:          value.Title,
			ExtId:          value.ExtId,
			CreatedAt:      value.CreatedAt,
			UpdatedAt:      value.UpdatedAt,
		}
		responses = append(responses, dto)
	}
	return responses
}

// FindByExtId implements ContentBlockService.
func (s *ContentBlockServiceImpl) FindByExtId(uuid string) dto.ContentBlockDTO {
	contentData, err := s.ContentBlockRepository.FindByExtId(uuid)
	if err != nil {
		panic(err)
	}
	contentResponse := dto.ContentBlockDTO{
		RawContent:     contentData.RawContent,
		RawContentType: contentData.RawContentType,
		Title:          contentData.Title,
		ExtId:          contentData.ExtId,
		CreatedAt:      contentData.CreatedAt,
		UpdatedAt:      contentData.UpdatedAt,
	}
	return contentResponse
}

// FindById implements ContentBlockService.
func (s *ContentBlockServiceImpl) FindById(id int) dto.ContentBlockDTO {
	contentData, err := s.ContentBlockRepository.FindById(id)
	if err != nil {
		panic(err)
	}
	contentResponse := dto.ContentBlockDTO{
		RawContent:     contentData.RawContent,
		RawContentType: contentData.RawContentType,
		Title:          contentData.Title,
		ExtId:          contentData.ExtId,
		CreatedAt:      contentData.CreatedAt,
		UpdatedAt:      contentData.UpdatedAt,
	}
	return contentResponse
}

// Update implements ContentBlockService.
func (s *ContentBlockServiceImpl) Update(id int, dto dto.UpsertContentBlockDTO) {
	contentModel := model.ContentBlock{
		RawContent:     dto.RawContent,
		RawContentType: dto.RawContentType,
		Title:          dto.Title,
	}
	s.ContentBlockRepository.Update(id, contentModel)
}
