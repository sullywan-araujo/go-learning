package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/gorillamux", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application-json")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Retorno HandleFunc GorillaMux"))
	}).Methods("Get", "Post")

	http.ListenAndServe(":8081", r) // 'r' meu handler, quem vai fazer o gerenciamento das minhas rotas;
}
