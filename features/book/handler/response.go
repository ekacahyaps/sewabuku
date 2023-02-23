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
	UserId      uint   `json:"user_id"`
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
		UserId:      dataCore.UserId,
	}
}
