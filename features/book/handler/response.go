package handler

import "sewabuku/features/book"

type BookResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Year        uint   `json:"year"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserId      uint   `json:"user_id"`
}

type AddBookResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Year        uint   `json:"year"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func AddBookToResponse(dataCore book.Core) AddBookResponse {
	return AddBookResponse{
		ID:          dataCore.ID,
		Title:       dataCore.Title,
		Author:      dataCore.Author,
		Year:        dataCore.Year,
		Price:       dataCore.Price,
		Description: dataCore.Description,
		Image:       dataCore.Image,
	}
}

type AllBooksResponse struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   uint   `json:"year"`
	Price  uint   `json:"price"`
	Image  string `json:"image"`
}

func AllBooksToResponse(dataCore book.Core) AllBooksResponse {
	return AllBooksResponse{
		ID:     dataCore.ID,
		Title:  dataCore.Title,
		Author: dataCore.Author,
		Year:   dataCore.Year,
		Price:  dataCore.Price,
		Image:  dataCore.Image,
	}
}

func ListAllBooksToResponse(dataCore []book.Core) []AllBooksResponse {
	var DataResponse []AllBooksResponse

	for _, value := range dataCore {
		DataResponse = append(DataResponse, AllBooksToResponse(value))
	}
	return DataResponse
}
