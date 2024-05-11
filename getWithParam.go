package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/:getname", func(c *gin.Context) {
		name := c.Param("getname")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	})

	router.Run(":8080")
}
