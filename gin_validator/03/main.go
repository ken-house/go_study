package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type ParamsData struct {
	Id       int       `json:"id" form:"id" binding:"required"`                                         // 必填
	Age      int       `json:"age" form:"age" binding:"required,gt=18"`                                 // 必填且值必须大于18 gt大于 gte大于等于 lt小于 let小于等于
	Salary   int       `json:"salary" form:"salary" binding:"omitempty,min=1000,max=10000"`             // 字段存在则进行验证值是否大于1000且小雨10000
	Birthday time.Time `json:"birthday" form:"birthday" time_format:"2006-01-02 15:04:05" time_utc:"0"` // 要求格式与2006-01-02 15:04:05一致
	Work     string    `json:"work" form:"work" binding:"-"`
	Gender   int       `json:"gender" form:"gender" binding:"oneof=0 1 2"`                  // 值只能是012中的一个
	Num      int       `json:"num" form:"num" binding:"ne=10"`                              // eq为等于，ne为不等于
	Sport    string    `json:"sport" form:"sport" binding:"len=3"`                          // 对字符串长度进行限制包括eq、ne、oneof、max、gt等
	Love     []string  `json:"love" form:"love" binding:"gt=0,dive,required,min=1,max=100"` // 对数组校验，dive前面的规则是对数组整体进行校验，dive后的规则是对元素进行校验
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
