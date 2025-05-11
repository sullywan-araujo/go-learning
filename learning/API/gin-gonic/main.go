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

	//QUERY STRINGS
	r.GET("/welcome", func(c *gin.Context) { // /welcome?firstname=Sullywan&lastname=Araujo&Guest=test
		firstName := c.DefaultQuery("firstname", "visitante")
		lastName := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Ola %s %s, seja bem vindo!!", firstName, lastName)
	})

	r.Run(":8080")
}
