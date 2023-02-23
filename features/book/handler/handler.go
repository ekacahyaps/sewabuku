package handler

import (
	"log"
	"mime/multipart"
	"net/http"
	"sewabuku/features/book"
	"sewabuku/helper"
	"strconv"

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
	return func(c echo.Context) error {

		res, err := bh.srv.AllBooks()
		if err != nil {
			log.Println("error running allproducts service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ListAllBooksToResponse(res),
			"message": "success show all products",
		})
	}
}

func (bh *bookHdl) MyBooks() echo.HandlerFunc {
	return nil
}

func (bh *bookHdl) BooksRented() echo.HandlerFunc {
	return nil
}

func (bh *bookHdl) BooksDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := c.Param("id")
		cnv, err := strconv.Atoi(input)
		if err != nil {
			log.Println("Books detail param error")
			return c.JSON(http.StatusBadRequest, "wrong product id")
		}

		res, err := bh.srv.BooksDetail(token, uint(cnv))
		if err != nil {
			log.Println("error running Book's detail service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    AddBookToResponse(res),
			"message": "success show book's detail",
		})
	}
}

func (bh *bookHdl) Update() echo.HandlerFunc {
	return nil
}

func (bh *bookHdl) Delete() echo.HandlerFunc {
	return nil
}
