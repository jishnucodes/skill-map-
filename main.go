package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my skill map!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
