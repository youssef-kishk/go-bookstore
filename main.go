package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/youssef-kishk/go-bookstore/pkg/config"
	"github.com/youssef-kishk/go-bookstore/pkg/controllers"
	controllerPort "github.com/youssef-kishk/go-bookstore/pkg/controllers"
	domainModels "github.com/youssef-kishk/go-bookstore/pkg/domain/model"
	domainService "github.com/youssef-kishk/go-bookstore/pkg/domain/service"
	repository "github.com/youssef-kishk/go-bookstore/pkg/repository/adapters"
)

// DB configuration params
const USERNAME = "root"
const PASS = "B211998."
const DATABASE = "bookstore"

var client *gorm.DB

func main() {
	conf, err := config.NewConfig(USERNAME, PASS, DATABASE)
	if err != nil {
		panic(err)
	}
	repo := repository.NewMysqlRepository(conf)
	controller := controllers.NewBookController(domainService.NewBookSerivce(repository.NewBookRepository(repo)))
	repo.Db.AutoMigrate(&domainModels.Book{})
	router := mux.NewRouter()
	configureRoutes(router, controller)

	http.Handle("/", router)
	fmt.Printf("Start server on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func configureRoutes(router *mux.Router, controller controllerPort.IBookController) {
	router.HandleFunc("/book", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book", controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
}
