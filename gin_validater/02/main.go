package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
)

type UserForm struct {
	Phone           string `json:"phone" form:"phone" binding:"required,len=11,validatePhone"`
	Password        string `json:"password" form:"password" binding:"required,max=10"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,eqfield=Password"`
}

var validatePhone validator.Func = func(fl validator.FieldLevel) bool {
	data := fl.Field().String()
	bool, err := regexp.MatchString("1[3-9][0-9]{9}", data)
	if err != nil {
		return false
	}
	return bool
}

func TestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params UserForm
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusOK, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	}
}

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validatePhone", validatePhone)
	}
	router := gin.Default()
	router.GET("/test", TestHandler())
	router.Run(":8086")
}
