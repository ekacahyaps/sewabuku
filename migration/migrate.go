package migration

import (
	user "sewabuku/features/user/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(user.User{})
}
