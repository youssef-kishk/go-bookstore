package domain

import (
	"errors"

	domainPort "github.com/youssef-kishk/go-bookstore/pkg/domain"
	domainModels "github.com/youssef-kishk/go-bookstore/pkg/domain/model"
	"github.com/youssef-kishk/go-bookstore/pkg/repository"
)

type BookSerivce struct {
	Repo repository.IBookRepository
}

func NewBookSerivce(repo repository.IBookRepository) domainPort.IBookService {
	service := &BookSerivce{repo}
	return service
}

func (service *BookSerivce) Create(b *domainModels.Book) (*domainModels.Book, error) {
	book := service.Repo.Create(b)
	return book, nil
}

func (service *BookSerivce) Update(id int64, updatedBook *domainModels.Book) (*domainModels.Book, error) {
	bookDetails, err := service.GetById(id)
	if err != nil {
		return bookDetails, err
	}
	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}
	service.Repo.Update(bookDetails)
	return bookDetails, nil
}

func (service *BookSerivce) GetAll() ([]domainModels.Book, error) {
	return service.Repo.GetAll(), nil
}

func (service *BookSerivce) GetById(id int64) (*domainModels.Book, error) {
	book := service.Repo.GetById(id)
	if book.ID == 0 {
		return book, errors.New("Book not found")
	}
	return book, nil
}

func (service *BookSerivce) Delete(id int64) (*domainModels.Book, error) {
	book, err := service.GetById(id)

	if err != nil {
		return book, err
	}
	return service.Repo.Delete(book), nil
}
