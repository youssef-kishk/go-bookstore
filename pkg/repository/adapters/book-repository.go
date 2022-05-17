package repository

import (
	domainModels "github.com/youssef-kishk/go-bookstore/pkg/domain/model"
	"github.com/youssef-kishk/go-bookstore/pkg/repository"
)

type BookRepository struct {
	baseRepo repository.BaseRepository
}

func NewBookRepository(base repository.BaseRepository) repository.IBookRepository {
	repo := &BookRepository{base}
	return repo

}

func (repo *BookRepository) Create(b *domainModels.Book) *domainModels.Book {
	repo.baseRepo.Db.NewRecord(&b)
	repo.baseRepo.Db.Create(&b)
	return b
}

func (repo *BookRepository) Update(b *domainModels.Book) *domainModels.Book {
	repo.baseRepo.Db.Save(&b)
	return b
}

func (repo *BookRepository) GetAll() []domainModels.Book {
	var books []domainModels.Book
	repo.baseRepo.Db.Find(&books)
	return books
}

func (repo *BookRepository) GetById(id int64) *domainModels.Book {
	book := &domainModels.Book{}
	repo.baseRepo.Db.Where("ID=?", id).Find(&book)
	return book
}

func (repo *BookRepository) Delete(book *domainModels.Book) *domainModels.Book {
	repo.baseRepo.Db.Delete(book)
	return book
}
