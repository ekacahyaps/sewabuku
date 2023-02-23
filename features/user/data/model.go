package data

import (
	"sewabuku/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	HP       string
	Email    string
	Address  string
	Photo    string
	Password string
}

func ToCore(data User) user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		HP:       data.HP,
		Email:    data.Email,
		Address:  data.Address,
		Photo:    data.Photo,
		Password: data.Password,
	}
}

func CoreToData(data user.Core) User {
	return User{
		Model:    gorm.Model{ID: data.ID},
		Name:     data.Name,
		HP:       data.HP,
		Email:    data.Email,
		Address:  data.Address,
		Photo:    data.Photo,
		Password: data.Password,
	}
}

func (dataModel *User) ModelsToCore() user.Core {
	return user.Core{
		ID:    dataModel.ID,
		Name:  dataModel.Name,
		Email: dataModel.Email,
	}
}
func listModelToCore(dataModel []User) []user.Core {
	var dataCore []user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ModelsToCore())
	}
	return dataCore
}
