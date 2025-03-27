package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//==========================================
	// simulando 'while';

	test := false

	for test == false {
		fmt.Println("VALOR FALSO")
	}

	//==========================================
	// simulando confirmação de pagamento

	pagamentoAprovado := false

	for !pagamentoAprovado {
		time.Sleep(2 * time.Second)

		pagamentoAprovado = processarPagamento()

		if !pagamentoAprovado {
			fmt.Println("Pagamento ainda não aprovado... tentando novamente.")
		}
	}

	fmt.Println("Pagamento aprovado. Liberando transação")
}

func processarPagamento() bool {
	return rand.Intn(5) == 0
}
