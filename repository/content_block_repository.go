package repository

import "github.com/LeonardoZ/gutti-headless-cms/model"

type ContentBlockRepository interface {
	Save(block model.ContentBlock) error
	Update(id int, block model.ContentBlock) error
	Delete(id int) error
	FindById(id int) (block *model.ContentBlock, err error)
	FindByExtId(uuid string) (block *model.ContentBlock, err error)
	FindAll() ([]model.ContentBlock, error)
}
