package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("dddddddddddddddddddddddddddddd")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":6010") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
