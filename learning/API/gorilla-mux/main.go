package main

import (
	"fmt"
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

	r.HandleFunc("/user/{name}/idade/{idade}", func(rw http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		nome := vars["name"]
		idade := vars["idade"]

		rw.Header().Set("Content-Type", "application-json")
		rw.WriteHeader(http.StatusOK)

		resp := fmt.Sprintf("Hello %s, voce tem %s anos", nome, idade)
		rw.Write([]byte(resp))
	})

	http.ListenAndServe(":8081", r) // 'r' meu handler, quem vai fazer o gerenciamento das minhas rotas;
}
