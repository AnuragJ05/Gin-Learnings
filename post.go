package main

import (
	"github.com/gin-gonic/gin"
)

type data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := gin.Default()

	var d1 data

	router.POST("/", func(c *gin.Context) {

		_ = c.BindJSON(&d1)

		c.JSON(200, gin.H{
			"name": d1.Name,
			"age":  d1.Age,
		})
	})

	router.Run(":8080")
}
