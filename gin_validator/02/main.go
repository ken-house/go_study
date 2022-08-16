package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type InnerUser struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required,max=10"`
}

type Param struct {
	InnerUserField  InnerUser
	ConfirmPassword string `json:"confirm_password" validate:"required,eqcsfield=InnerUserField.Password"`
}

func main() {
	user := InnerUser{
		Username: "lisi",
		Password: "123456",
	}

	param := Param{
		InnerUserField:  user,
		ConfirmPassword: "12345",
	}
	var validate = validator.New()
	errs := validate.Struct(param)
	fmt.Println(errs)
}
