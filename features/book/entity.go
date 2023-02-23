package book

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	Title       string
	Author      string
	Year        uint
	Price       uint
	Description string
	Image       string
	UserId      uint
}

type BookHandler interface {
	Add() echo.HandlerFunc
	AllBooks() echo.HandlerFunc
	MyBooks() echo.HandlerFunc
	BooksRented() echo.HandlerFunc
	BooksDetail() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type BookService interface {
	Add(token interface{}, newBook Core, image *multipart.FileHeader) (Core, error)
	AllBooks() ([]Core, error)
	MyBooks(token interface{}) ([]Core, error)
	BooksRented(token interface{}) ([]Core, error)
	BooksDetail(token interface{}, bookId uint) (Core, error)
	Update(token interface{}, bookId uint, updateBook Core, updateImage *multipart.FileHeader) (Core, error)
	Delete(token interface{}, bookId uint) error
}

type BookData interface {
	Add(userId uint, newBook Core) (Core, error)
	AllBooks() ([]Core, error)
	MyBooks(userId uint) ([]Core, error)
	BooksRented(userId uint) ([]Core, error)
	BooksDetail(userId, bookId uint) (Core, error)
	Update(userId uint, bookId uint, updateBook Core) (Core, error)
	Delete(userId, bookId uint) error
}
