package data

import (
	"sewabuku/features/book"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Author      string
	Year        uint
	Price       uint
	Description string
	Image       string
	UserId      uint
}

func DataToCore(data Book) book.Core {
	return book.Core{
		ID:          data.ID,
		Title:       data.Title,
		Author:      data.Author,
		Year:        data.Year,
		Price:       data.Price,
		Description: data.Description,
		Image:       data.Image,
		UserId:      data.UserId,
	}
}

func CoreToData(data book.Core) Book {
	return Book{
		Model:       gorm.Model{ID: data.ID},
		Title:       data.Title,
		Author:      data.Author,
		Year:        data.Year,
		Price:       data.Price,
		Description: data.Description,
		Image:       data.Image,
		UserId:      data.UserId,
	}
}

type AllBook struct {
	ID          uint
	Title       string
	Author      string
	Year        uint
	Price       uint
	Description string
	Image       string
}

func (dataModel *AllBook) AllBooksToCore() book.Core {
	return book.Core{
		ID:          dataModel.ID,
		Title:       dataModel.Title,
		Author:      dataModel.Author,
		Year:        dataModel.Year,
		Price:       dataModel.Price,
		Description: dataModel.Description,
		Image:       dataModel.Image,
	}
}

func ListAllBooksToCore(dataModels []AllBook) []book.Core {
	var dataCore []book.Core
	for _, value := range dataModels {
		dataCore = append(dataCore, value.AllBooksToCore())
	}
	return dataCore
}
