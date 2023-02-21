package data

import (
	"errors"
	"log"
	"sewabuku/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Login(email string) (user.Core, error) {
	userLogin := User{}

	if err := uq.db.Where("email = ?", email).First(&userLogin).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return ToCore(userLogin), nil
}
func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {

	cnv := CoreToData(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		return user.Core{}, err
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) Profile(id uint) (user.Core, error) {
	return user.Core{}, nil
}

func (uq *userQuery) Update(id uint, updatedData user.Core) (user.Core, error) {
	return user.Core{}, nil
}

func (uq *userQuery) Delete(id uint) error {
	return nil
}
