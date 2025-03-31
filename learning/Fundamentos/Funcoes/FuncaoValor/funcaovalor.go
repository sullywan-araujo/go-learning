package funcaovalor

import "fmt"

func Saque(valor float64) {
	fmt.Printf("Saque de %.2f realizado com sucesso.\n", valor)
}

func Deposito(valor float64) {
	fmt.Printf("Depósito de %.2f realizado com sucesso.\n", valor)
}

func ProcessarOperacao(valor float64, operacao func(float64)) {
	operacao(valor)
}
