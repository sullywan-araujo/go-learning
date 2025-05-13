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

	// QUERY STRING
	r.HandleFunc("/transacoes", BuscarTransacaoHandler).Methods("GET")

	http.ListenAndServe(":8081", r) // 'r' meu handler, quem vai fazer o gerenciamento das minhas rotas;
}

func BuscarTransacaoHandler(rw http.ResponseWriter, req *http.Request) {
	tipoTransacao := req.URL.Query().Get("tipo")
	mesTransacao := req.URL.Query().Get("mes")

	var tipoDescricao string
	var mesDescricao string

	switch tipoTransacao {
	case "pix":
		tipoDescricao = "Pix"
	case "cartaoCredito":
		tipoDescricao = "Cartão de Crédito"
	case "cartaoDebito":
		tipoDescricao = "Cartão de Débito"
	default:
		tipoDescricao = "Tipo de transação inválido ou não informado"
	}

	switch mesTransacao {
	case "1":
		mesDescricao = "Janeiro"
	case "2":
		mesDescricao = "Fevereiro"
	case "3":
		mesDescricao = "Março"
	case "4":
		mesDescricao = "Abril"
	case "5":
		mesDescricao = "Maio"
	case "6":
		mesDescricao = "Junho"
	case "7":
		mesDescricao = "Julho"
	case "8":
		mesDescricao = "Agosto"
	case "9":
		mesDescricao = "Setembro"
	case "10":
		mesDescricao = "Outubro"
	case "11":
		mesDescricao = "Novembro"
	case "12":
		mesDescricao = "Dezembro"
	default:
		mesDescricao = "Mês inválido ou não informado"
	}

	resposta := fmt.Sprintf("Tipo de transação: %s\nMês: %s", tipoDescricao, mesDescricao)

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(resposta))

}
