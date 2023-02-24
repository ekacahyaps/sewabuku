package cart

import "github.com/labstack/echo/v4"

type Core struct {
	ID         uint
	TotalPrice int
	BookId     uint
	Title      string
	UserID     uint
}

type CartHandler interface {
	AddCart() echo.HandlerFunc
	MyCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
}

type CartService interface {
	AddCartSrv(token interface{}, newCart Core) (Core, error)
	MyCartSrv(token interface{}) ([]Core, error)
	DeleteCartSrv(token interface{}, cartID uint) error
}

type CartData interface {
	AddCartData(userID uint, newCart Core) (Core, error)
	AddCHPData(cartID uint, newCart Core) (Core, error)
	MyCartData(id uint) ([]Core, error)
	DeleteCartData(userID, cartID uint) error
}
