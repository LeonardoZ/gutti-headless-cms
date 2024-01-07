package service

import (
	"net/http"

	"github.com/LeonardoZ/gutti-headless-cms/dto"
	"github.com/LeonardoZ/gutti-headless-cms/error_handler"
	"github.com/LeonardoZ/gutti-headless-cms/model"
	"github.com/LeonardoZ/gutti-headless-cms/repository"
)

type ContentBlockServiceImpl struct {
	ContentBlockRepository repository.ContentBlockRepository
}

func NewContentBlockService(contentRepository repository.ContentBlockRepository) ContentBlockService {
	return &ContentBlockServiceImpl{
		ContentBlockRepository: contentRepository,
	}
}

// Create implements ContentBlockService.
func (s *ContentBlockServiceImpl) Create(dto dto.UpsertContentBlockDTO) error {
	model := model.ContentBlock{
		Title:          dto.Title,
		RawContent:     dto.RawContent,
		RawContentType: dto.RawContentType,
	}

	err := s.ContentBlockRepository.Save(model)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements ContentBlockService.
func (s *ContentBlockServiceImpl) Delete(id int) error {
	found, err := s.FindById(id)
	if err != nil {
		return err
	}

	if found == nil {
		return error_handler.NewHttpError("Failed to fetch content by id", "", http.StatusNotFound)
	}
	err = s.ContentBlockRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements ContentBlockService.
func (s *ContentBlockServiceImpl) FindAll() ([]dto.ContentBlockDTO, error) {
	result, err := s.ContentBlockRepository.FindAll()
	if err != nil {
		return nil, err
	}
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
	return responses, nil
}

// FindByExtId implements ContentBlockService.
func (s *ContentBlockServiceImpl) FindByExtId(uuid string) (*dto.ContentBlockDTO, error) {
	contentData, err := s.ContentBlockRepository.FindByExtId(uuid)
	if err != nil {
		return nil, err
	}
	if contentData == nil {
		return nil, error_handler.NewHttpError("Failed to fetch content by id", "", http.StatusNotFound)
	}
	contentResponse := dto.ContentBlockDTO{
		RawContent:     contentData.RawContent,
		RawContentType: contentData.RawContentType,
		Title:          contentData.Title,
		ExtId:          contentData.ExtId,
		CreatedAt:      contentData.CreatedAt,
		UpdatedAt:      contentData.UpdatedAt,
	}
	return &contentResponse, nil
}

// FindById implements ContentBlockService.
func (s *ContentBlockServiceImpl) FindById(id int) (*dto.ContentBlockDTO, error) {
	contentData, err := s.ContentBlockRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if contentData == nil {
		return nil, error_handler.NewHttpError("Failed to fetch content by id", "", http.StatusNotFound)
	}
	contentResponse := dto.ContentBlockDTO{
		RawContent:     contentData.RawContent,
		RawContentType: contentData.RawContentType,
		Title:          contentData.Title,
		ExtId:          contentData.ExtId,
		CreatedAt:      contentData.CreatedAt,
		UpdatedAt:      contentData.UpdatedAt,
	}
	return &contentResponse, nil
}

// Update implements ContentBlockService.
func (s *ContentBlockServiceImpl) Update(id int, dto dto.UpsertContentBlockDTO) error {
	found, err := s.FindById(id)
	if err != nil {
		return err
	}

	if found == nil {
		return error_handler.NewHttpError("Failed to fetch content by id", "", http.StatusNotFound)
	}
	contentModel := model.ContentBlock{
		RawContent:     dto.RawContent,
		RawContentType: dto.RawContentType,
		Title:          dto.Title,
	}
	err = s.ContentBlockRepository.Update(id, contentModel)
	if err != nil {
		return err
	}
	return nil
}
