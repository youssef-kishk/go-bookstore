package domain

import domainModels "github.com/youssef-kishk/go-bookstore/pkg/domain/model"

type IBookService interface {
	Create(b *domainModels.Book) (*domainModels.Book, error)
	Update(id int64, b *domainModels.Book) (*domainModels.Book, error)
	GetAll() ([]domainModels.Book, error)
	GetById(id int64) (*domainModels.Book, error)
	Delete(id int64) (*domainModels.Book, error)
}
