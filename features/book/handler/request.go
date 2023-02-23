package handler

import "sewabuku/features/book"

type AddBookRequest struct {
	Title       string `form:"title"`
	Author      string `form:"author"`
	Year        uint   `form:"year"`
	Price       uint   `form:"price"`
	Description string `form:"description"`
	Image       string `form:"image"`
}

func ToCore(data interface{}) *book.Core {
	res := book.Core{}

	switch data.(type) {
	case AddBookRequest:
		cnv := data.(AddBookRequest)
		res.Title = cnv.Title
		res.Author = cnv.Author
		res.Year = cnv.Year
		res.Price = cnv.Price
		res.Description = cnv.Description
		res.Image = cnv.Image
	default:
		return nil
	}

	return &res
}
