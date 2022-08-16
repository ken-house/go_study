package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ParamsData struct {
	Hobby []string `json:"hobby" form:"hobby" validate:"required,unique"`
	Love  []string `json:"love" form:"love" binding:"gt=0,dive,required,min=1,max=100"` // 对数组校验，dive前面的规则是对数组整体进行校验，dive后的规则是对元素进行校验
}

func main() {
	param := ParamsData{Hobby: []string{
		"swimming",
		"football",
		"football",
	}}
	var validate = validator.New()
	err := validate.Struct(param)
	fmt.Println(err)
}
