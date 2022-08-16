package main

import "github.com/gin-gonic/gin"

type ParamsData struct {
	Username string `json:"username" form:"username" binding:"contains=123456"`
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
