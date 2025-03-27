package FuncaoSemRetorno

import "fmt"

func ExibirSaldo(nome string, saldo float64) {
	fmt.Printf("Cliente: %s | Saldo atual: R$ %2.f\n", nome, saldo)
}
