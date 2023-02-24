package service

import (
	"errors"
	"mime/multipart"
	"sewabuku/features/book"
	"sewabuku/helper"
	"strings"
)

type bookSrv struct {
	data book.BookData
}

func New(d book.BookData) book.BookService {
	return &bookSrv{
		data: d,
	}
}

func (bs *bookSrv) Add(token interface{}, newBook book.Core, image *multipart.FileHeader) (book.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return book.Core{}, errors.New("user not found")
	}

	if image != nil {
		path, err := helper.UploadProductImageS3(*image, helper.ExtractToken(token))
		if err != nil {
			return book.Core{}, err
		}
		newBook.Image = path
	}

	res, err := bs.data.Add(uint(userID), newBook)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book not found"
		} else {
			msg = "server problem"
		}
		return book.Core{}, errors.New(msg)
	}

	return res, nil
}

func (bs *bookSrv) AllBooks() ([]book.Core, error) {
	res, err := bs.data.AllBooks()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return []book.Core{}, errors.New(msg)
	}
	return res, nil
}

func (bs *bookSrv) MyBooks(token interface{}) ([]book.Core, error) {
	userId := helper.ExtractToken(token)
	res, err := bs.data.MyBooks(uint(userId))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return []book.Core{}, errors.New(msg)
	}
	return res, nil
}

func (bs *bookSrv) BooksRented(token interface{}) ([]book.Core, error) {
	return []book.Core{}, nil
}

func (bs *bookSrv) BooksDetail(token interface{}, bookId uint) (book.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := bs.data.BooksDetail(uint(userID), bookId)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return book.Core{}, errors.New(msg)
	}
	return res, nil
}

func (bs *bookSrv) Update(token interface{}, bookId uint, updateBook book.Core, updateImage *multipart.FileHeader) (book.Core, error) {
	return book.Core{}, nil
}

func (bs *bookSrv) Delete(token interface{}, bookId uint) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("user not found")
	}

	err := bs.data.Delete(uint(userID), bookId)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book not found"
		} else {
			msg = "server problem"
		}
		return errors.New(msg)
	}
	return nil
}
