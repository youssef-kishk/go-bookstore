package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/youssef-kishk/go-bookstore/pkg/config"
	"github.com/youssef-kishk/go-bookstore/pkg/repository"
)

func NewMysqlRepository(conf *config.Config) repository.BaseRepository {

	uri := fmt.Sprintf("%s:%s@tcp(db)/%s?charset=utf8&parseTime=True&loc=Local", conf.USERNAME, conf.PASS, conf.DBNAME)
	d, err := gorm.Open("mysql", uri)
	if err != nil {
		panic(err)
	}
	repo := repository.BaseRepository{d}
	return repo
}
