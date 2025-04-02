package funcaometodo

import "fmt"

type ContaBancaria struct {
	Titular string
	Saldo   float64
}

func (c *ContaBancaria) Depositar(valor float64) {
	if valor <= 0 {
		fmt.Println("O valor do depósito deve ser maior que zero.")
		return
	}

	c.Saldo += valor
	fmt.Printf("Depósito de %.2f realizado para %s. Saldo atualizado é de R$ %.2f\n", valor, c.Titular, c.Saldo)
}
