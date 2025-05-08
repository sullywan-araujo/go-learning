package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/gochi", func(rw http.ResponseWriter, req *http.Request) {
		// a função write é responsável por converter a string em bytes e envia os bytes pela conexão HTTP.
		rw.Write([]byte("Retornando requisição com go-chi")) // Em Go, tudo que é enviado como resposta HTTP precisa estar em formato binário (bytes)
	})

	http.ListenAndServe(":8080", r)
}
