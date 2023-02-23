package migration

import (
	book "sewabuku/features/book/data"
	user "sewabuku/features/user/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(user.User{})
	db.AutoMigrate(book.Book{})
}
