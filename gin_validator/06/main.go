package main

import "github.com/gin-gonic/gin"

type ParamsData struct {
	Age    int    `json:"age" form:"age" binding:"required"`
	Gender int    `json:"gender" form:"gender" binding:"required,oneof=1 2"`
	Name   string `json:"name" form:"name" binding:"required_if=Age 18 Gender 1"`
}

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params ParamsData
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
