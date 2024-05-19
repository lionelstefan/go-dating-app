package request

import (
	"github.com/lionelstefan/go-dating-app/model"
)

type RegisterRequest struct {
	Email string `json:email`
	Password string `json:password`
	FullName string `json:fullname`
	Age string `json:age`
	Gender string `json:gender`
}

func (r *RegisterRequest) ToModel() (result model.UserRegister) {
	result = model.UserRegister{
		Email: r.Email,
		Password: r.Password,
		FullName: r.FullName,
		Age: r.Age,
		Gender: r.Gender,
		IsVerified: false,
		NoSwipeLimit: false,
	}

	return
}