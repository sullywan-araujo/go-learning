package FuncaoComRetorno

import "fmt"

// Calcula saldo após saque
func CalcularSaldo(saldo, saque float64) float64 {
	if saldo > saque {
		return saldo - saque
	}
	fmt.Printf("Saldo insuficiente: %2.f\n", saldo)
	return saldo
}
