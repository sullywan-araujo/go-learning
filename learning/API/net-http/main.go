package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Para implementarmos esse pacote é necessário startar o handler;
	// Handler = atua como um manipulador de requisições;
	// Handle = registra um handler, realiza o direcionamento da URI
	// HandleFunc = faz a conversão da função para um handler

	http.Handle("/apihandler", MeuHandler{}) //localhost:8080/apihandler

	http.HandleFunc("/apihandlefunc", func(w http.ResponseWriter, r *http.Request) { //localhost:8080/apihandlefunc
		fmt.Fprintln(w, "Resposta do HandleFunc")
	})

	fmt.Println("Servidor localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type MeuHandler struct{}

func (h MeuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, MeuHandler!"))
}

func Handle(url string, handler MeuHandler) {
}
