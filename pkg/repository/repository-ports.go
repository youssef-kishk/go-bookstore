package repository

import (
	"github.com/jinzhu/gorm"
	domainModels "github.com/youssef-kishk/go-bookstore/pkg/domain/model"
)

type BaseRepository struct {
	Db *gorm.DB
}

type IBookRepository interface {
	Create(b *domainModels.Book) *domainModels.Book
	Update(b *domainModels.Book) *domainModels.Book
	GetAll() []domainModels.Book
	GetById(id int64) *domainModels.Book
	Delete(book *domainModels.Book) *domainModels.Book
}
