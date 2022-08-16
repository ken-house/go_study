package main

import "github.com/gin-gonic/gin"

type ParamData struct {
	Password        string `json:"password" form:"password" binding:"required,max=10"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,eqfield=Password"`
}

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params ParamData
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"params": params,
		})
	}
}

func main() {
	router := gin.Default()
	router.POST("/test", Test())
	router.Run(":8082")
}
