package data

import (
	"log"
	"sewabuku/features/book"

	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.BookData {
	return &bookData{
		db: db,
	}
}

func (bd *bookData) Add(userID uint, newBook book.Core) (book.Core, error) {
	convert := CoreToData(newBook)
	convert.UserId = userID

	err := bd.db.Create(&convert).Error
	if err != nil {
		log.Println("add book query error", err.Error())
		return book.Core{}, err
	}

	newBook.ID = convert.ID
	newBook.UserId = convert.UserId

	return newBook, nil
}

func (bd *bookData) AllBooks() ([]book.Core, error) {
	return []book.Core{}, nil
}

func (bd *bookData) MyBooks(userId uint) ([]book.Core, error) {
	return []book.Core{}, nil
}

func (bd *bookData) BooksRented(userId uint) ([]book.Core, error) {
	return []book.Core{}, nil
}

func (bd *bookData) BooksDetail(userId, bookId uint) (book.Core, error) {
	return book.Core{}, nil
}

func (bd *bookData) Update(userId uint, bookId uint, updateBook book.Core) (book.Core, error) {
	return book.Core{}, nil
}

func (bd *bookData) Delete(userId, bookId uint) error {
	return nil
}
