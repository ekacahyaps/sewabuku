package handler

import "sewabuku/features/user"

type UserReponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	HP      string `json:"hp"`
	Address string `json:"address"`
	Photo   string `json:"photo"`
}

type RegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		Name:    data.Name,
		Email:   data.Email,
		HP:      data.HP,
		Address: data.Address,
		Photo:   data.Photo,
	}
}

func ToResponses(data user.Core) RegisterResponse {
	return RegisterResponse{
		Name:  data.Name,
		Email: data.Email,
	}
}
func fromCoreList(dataCore []user.Core) []UserReponse {
	var dataResponse []UserReponse

	for _, v := range dataCore {
		dataResponse = append(dataResponse, ToResponse(v))
	}
	return dataResponse
}
