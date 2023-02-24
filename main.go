package main

import (
	"sewabuku/config"
	"sewabuku/features/user/data"
	"sewabuku/features/user/handler"
	services "sewabuku/features/user/service"
	"sewabuku/migration"

	bdata "sewabuku/features/book/data"
	bhdl "sewabuku/features/book/handler"
	bsrv "sewabuku/features/book/service"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	migration.Migrate(db)

	userData := data.New(db)
	userSrv := services.New(userData)
	userHdl := handler.New(userSrv)

	bookData := bdata.New(db)
	bookSrv := bsrv.New(bookData)
	bookHdl := bhdl.New(bookSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/books", bookHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/books", bookHdl.AllBooks())
	e.GET("/books/:id", bookHdl.BooksDetail(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/books/:id", bookHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/mybooks", bookHdl.MyBooks(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
