package handler

import "sewabuku/features/user"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Hp       string `json:"hp" form:"hp"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}

func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Email = cnv.Email
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Name = cnv.Name
		res.HP = cnv.Hp
		res.Email = cnv.Email
		res.Address = cnv.Address
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}
