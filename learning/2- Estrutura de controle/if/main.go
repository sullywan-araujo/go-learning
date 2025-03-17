package main

import "fmt"

func main() {
	//IF
	if 2 > 1 {
		fmt.Println("true")
	}

	// IF / ELSE
	if 1 > 1 {
		fmt.Println("Equals")
	} else {
		fmt.Println("Different")
	}

	// IF / ELSE IF
	if 2 < 1 {
		fmt.Println("falso")
	} else if 5 < 1 {
		fmt.Println("falso2")
	} else {
		fmt.Println("verdadeiro")
	}

	// INSTANCIANDO VARIAVEL NO IF
	if retornoTest := test(); retornoTest == "qlqrCoisa" { // a variavel instanciada não vai funcionar fora do 'if'
		fmt.Println("retorno verdadeiro")
	}
}

func test() string {
	return "qlqrCoisa"
}
