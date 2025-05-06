package main

import "fmt"

func main() {
	Print("Olá")
	Print(42)

	fmt.Println(Soma(10, 20))
	fmt.Println(Soma(3.75, 4.56))

	caixa := Caixa[string]{Valor: "Texto dentro da caixa"}
	fmt.Println("Caixa:", caixa.Get())

	ok := Contem([]int{1, 2, 3}, 2)
	fmt.Println("Contem o 2?", ok)
}
