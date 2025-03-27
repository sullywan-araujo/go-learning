package main

import "fmt"

func main() {
	// simulando do-while
	anExpression := false

	for ok := true; ok; ok = anExpression { // como estamos simulando um do-while, o código vai executar uma vez, e depois vai realizar a verificação para dar continuidade no código;
		fmt.Println("Passou por aqui")
	}

	//==================================================
	// simulando tentativa de saque

	saldo := 100.00
	valorSaque := 50.00
	autorizado := true

	for ok := true; ok; ok = autorizado {
		fmt.Printf("Tentativa de saque de R$ %2.f\n", valorSaque)

		if saldo >= valorSaque {
			saldo -= valorSaque

			fmt.Println("Saque realizado com sucesso.")
			fmt.Printf("Seu saldo é de: R$ %2.f\n", saldo)
		} else {
			fmt.Println("Saque negado, saldo insuficiente.")
			autorizado = false
		}

	}
}
