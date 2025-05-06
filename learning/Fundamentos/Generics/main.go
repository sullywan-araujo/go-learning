package main

import (
	"fmt"

	tr "github.com/learning/Fundamentos/Generics/Transacao"
)

// nem todo tipo 'comparable' com == e != tambem pode ser comparado com operadores como >, <, >=, <=;
// 'Comparable' na restrição genérica só garante que os tipos podem ser usados com == e !=, não com >, <;
//func Maior[T comparable](a, b T) T {
//	if a > b {
//		return a
//	}
//	return b
//}

type Ordenavel interface {
	~int | ~float64 | ~string
}

// T é um parametro do tipo generico;
func MaioCerto[T Ordenavel](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Soma[T int | float64](valores []T) T {
	var total T

	for _, v := range valores {
		total += v
	}

	return total
}

func main() {
	fmt.Println(MaioCerto("bola", "futebol"))

	fmt.Println(Soma([]float64{2.5, 2.5, 5}))

	tx1 := tr.Transacao[float64]{ID: "TX1001", Valor: 150.00}
	tx2 := tr.Transacao[float64]{ID: "TX1002", Valor: 1550.00}

	maior := tr.TransacaoMaiorValor(tx1, tx2)
	fmt.Printf("Maior transação ID=%s, Valor= R$ %.2f\n", maior.ID, maior.Valor)

	tx3 := tr.Transacao[int]{ID: "TX1003", Valor: 549}
	tx4 := tr.Transacao[int]{ID: "TX1004", Valor: 1549}
	maiorInt := tr.TransacaoMaiorValor(tx3, tx4)
	fmt.Printf("Maior transação ID=%s, Valor= R$ %d\n", maior.ID, maiorInt.Valor)

}

/*
Type Parameters (Parâmetros de Tipo) – São os nomes que você usa para se referir aos tipos genéricos (como T).

Type Constraints (Restrições de Tipo) – Dizem o que os tipos genéricos precisam suportar. Ex: comparable, any, ou uma interface com métodos.

Instantiation – Quando você chama a função, o compilador substitui T pelo tipo real que você passou.
*/
