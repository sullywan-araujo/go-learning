package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/gin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "gonic",
		})
	})

	// PARAMETROS NA ROTA
	r.GET("/user/:nome/idade/:idade", func(c *gin.Context) {
		nome := c.Param("name")
		idade := c.Param("idade")
		c.String(http.StatusOK, "Olá %s, voce tem %s anos", nome, idade)
	})
	r.Run(":8080")
}
