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
func (r *ContentBlockRepositoryImpl) Delete(id int) {
	var block model.ContentBlock
	result := r.Db.Where("id = ?", id).Delete(&block)
	if result.Error != nil {
		panic(result.Error)
	}
}

// FindAll implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) FindAll() []model.ContentBlock {
	var blocks []model.ContentBlock
	result := r.Db.Find(&blocks)
	if result.Error != nil {
		panic(result.Error)
	}
	return blocks
}

// FindByExtId implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) FindByExtId(extId string) (model.ContentBlock, error) {
	var block model.ContentBlock
	result := r.Db.Find(&block, extId)
	if result != nil {
		return block, nil
	} else {
		return block, errors.New("Content block not found with external id")
	}
}

// FindById implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) FindById(id int) (model.ContentBlock, error) {
	var block model.ContentBlock
	result := r.Db.Find(&block, id)
	if result != nil {
		return block, nil
	} else {
		return block, errors.New("Content block not found")
	}
}

// Save implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) Save(block model.ContentBlock) {
	result := r.Db.Omit("ext_id", "created_at", "updated_at").Create(&block)
	if result.Error != nil {
		panic(result.Error)
	}
}

// Update implements ContentBlockRepository.
func (r *ContentBlockRepositoryImpl) Update(id int, block model.ContentBlock) {
	found, err := r.FindById(id)
	if err != nil {
		panic(err)
	}
	result := r.Db.Model(&found).Updates(block)
	if result.Error != nil {
		panic(err)
	}
}
