package repository

import (
	"errors"

	"github.com/LeonardoZ/gutti-headless-cms/model"
	"gorm.io/gorm"
)

type ContentBlockRepositoryImpl struct {
	Db *gorm.DB
}

func NewContentBlockRepository(Db *gorm.DB) ContentBlockRepository {
	return &ContentBlockRepositoryImpl{Db: Db}
}

// Delete implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) Delete(id int) error {
	var block model.ContentBlock
	result := r.Db.Where("id = ?", id).Delete(&block)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindAll implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) FindAll() ([]model.ContentBlock, error) {
	var blocks []model.ContentBlock
	result := r.Db.Find(&blocks)
	if result.Error != nil {
		return nil, result.Error
	}
	return blocks, nil
}

// FindByExtId implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) FindByExtId(extId string) (*model.ContentBlock, error) {
	var block model.ContentBlock
	result := r.Db.Find(&block, extId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result != nil {
		return &block, nil
	} else {
		return &block, errors.New("Content block not found with external id")
	}
}

// FindById implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) FindById(id int) (*model.ContentBlock, error) {
	var block model.ContentBlock
	err := r.Db.First(&block, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &block, nil
}

// Save implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) Save(block model.ContentBlock) error {
	result := r.Db.Omit("ext_id", "created_at", "updated_at").Create(&block)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) Update(id int, block model.ContentBlock) error {
	found, err := r.FindById(id)
	if err != nil {
		return err
	}
	if found == nil {
		return errors.New("No content to update")
	}
	result := r.Db.Model(&found).Updates(block)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
