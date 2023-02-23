package handler

import (
	"log"
	"mime/multipart"
	"net/http"
	"sewabuku/features/book"
	"sewabuku/helper"

	"github.com/labstack/echo/v4"
)

type bookHdl struct {
	srv book.BookService
}

func New(bs book.BookService) book.BookHandler {
	return &bookHdl{
		srv: bs,
	}
}

func (bh *bookHdl) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := AddBookRequest{}
		var image *multipart.FileHeader
		if err := c.Bind(&input); err != nil {
			log.Println("add book req body scan error")
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}

		file, err := c.FormFile("image")
		if file != nil && err == nil {
			image = file
		} else if file != nil && err != nil {
			log.Println("error read product image")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		}

		res, err := bh.srv.Add(token, *ToCore(input), image)
		if err != nil {
			log.Println("error running add product service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    AddBookToResponse(res),
			"message": "success add book",
		})
	}
}

func (bh *bookHdl) AllBooks() echo.HandlerFunc {
	return nil
}

func (bh *bookHdl) MyBooks() echo.HandlerFunc {
	return nil
}

func (bh *bookHdl) BooksRented() echo.HandlerFunc {
	return nil
}

func (bh *bookHdl) BooksDetail() echo.HandlerFunc {
	return nil
}

func (bh *bookHdl) Update() echo.HandlerFunc {
	return nil
}

func (bh *bookHdl) Delete() echo.HandlerFunc {
	return nil
}
