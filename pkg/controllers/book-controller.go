package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	domainPort "github.com/youssef-kishk/go-bookstore/pkg/domain"
	domain "github.com/youssef-kishk/go-bookstore/pkg/domain/model"
	"github.com/youssef-kishk/go-bookstore/pkg/utils"
)

type BookController struct {
	service domainPort.IBookService
}

func NewBookController(service domainPort.IBookService) IBookController {
	controller := &BookController{service}
	return controller
}

func (c BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(books)
	utils.WriteHeader(w, http.StatusOK, res)
}

func (c BookController) GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}
	bookDetails, err := c.service.GetById(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(bookDetails)
	utils.WriteHeader(w, http.StatusOK, res)
}

func (c BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &domain.Book{}
	utils.ParseBody(r, book)
	b, err := c.service.Create(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(b)
	utils.WriteHeader(w, http.StatusOK, res)
}

func (c BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	_, err = c.service.Delete(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	utils.WriteHeader(w, http.StatusOK)
}

func (c BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := &domain.Book{}
	utils.ParseBody(r, updatedBook)
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}
	bookDetails, err := c.service.Update(ID, updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(bookDetails)
	utils.WriteHeader(w, http.StatusOK, res)
}
