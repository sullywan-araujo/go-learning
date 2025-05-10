package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// Para implementarmos esse pacote é necessário startar o handler;
	// Handler = atua como um manipulador de requisições; realiza o tratamento dos endpoints;
	// Handle = registra um handler, realiza o direcionamento da requisição
	// HandleFunc = faz a conversão da função para um handler

	http.Handle("/apihandler", MeuHandler{}) //localhost:8080/apihandler

	http.Handle("/user/", Handler2{})

	http.HandleFunc("/apihandlefunc", func(w http.ResponseWriter, r *http.Request) { //localhost:8080/apihandlefunc
		fmt.Fprintln(w, "Resposta do HandleFunc")
	})

	fmt.Println("Servidor localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type MeuHandler struct{}

func (h MeuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Olá, MeuHandler!"))

}

//==================================================

type Handler struct{}

func (hn Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// Extraindo o nome da URL;
	name := strings.TrimPrefix(req.URL.Path, "/user/")
	response := fmt.Sprintf("Hello %s", name)

	rw.Header().Set("Content-Type", "application-json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(response))
}

// ======================================================
// Extraindo mais de 1 parâmetro
type Handler2 struct{}

func (hd Handler2) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	parts := strings.Split(req.URL.Path, "/") //parsing manual
	// [ "", "user", "joao", "idade", "30" ]
	if len(parts) < 5 {
		http.Error(rw, "Parâmetros incompletos", http.StatusBadRequest)
	}

	name := parts[2]
	age := parts[4]

	response := fmt.Sprintf("Ola %s, voce tem %s anos", name, age)

	rw.Header().Set("Content-Type", "application-json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(response))
}
