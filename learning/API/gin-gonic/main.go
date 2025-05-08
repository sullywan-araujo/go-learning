package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/gin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "gonic",
		})
	})
	r.Run(":8080")
}
