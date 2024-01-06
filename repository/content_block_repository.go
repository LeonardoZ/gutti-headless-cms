package repository

import "github.com/LeonardoZ/gutti-headless-cms/model"

type ContentBlockRepository interface {
	Save(block model.ContentBlock)
	Update(id int, block model.ContentBlock)
	Delete(id int)
	FindById(id int) (block model.ContentBlock, err error)
	FindByExtId(uuid string) (block model.ContentBlock, err error)
	FindAll() []model.ContentBlock
}
