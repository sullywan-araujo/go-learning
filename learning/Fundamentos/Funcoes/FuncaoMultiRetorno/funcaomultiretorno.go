package funcaomultiretorno

import "fmt"

func Sacar(saldo, valorSaque float64) (float64, bool) {
	if saldo > valorSaque {
		return saldo - valorSaque, true
	}

	fmt.Println("Saldo insuficiente")
	return saldo, false
}
