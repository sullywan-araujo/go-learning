package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"idade" validate:"required,min=0,max=150"`
	Doc  string `json:"cpf" validate:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/welcome", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, "Erro ao converter dados")
			return
		}

		validate := validator.New()
		err := validate.Struct(user)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, "Erro ao validar dados")
			return
		}

		c.JSON(http.StatusOK, user)

	})

	fmt.Println("Servidor rodando na porta 8080")
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
