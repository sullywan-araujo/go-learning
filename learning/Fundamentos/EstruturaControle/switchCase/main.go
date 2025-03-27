package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// inicializando variável diretamente no switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MAC OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", os)
	}

	//==============================================
	// casos declarados
	// importante lembrar de sempre verificar os tipos de dados comparados

	b := "1"
	switch b {
	case "1":
		fmt.Println("um")
	case "2":
		fmt.Println("dois")
	}

	// neste caso nem vai compilar o código, pois os tipos comparados são diferentes;

	//==============================================
	// switch sem clausula de comparação
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 17:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}

	//==============================================
	// case chamando uma função

	a := 33

	switch {
	case parMenor10(a):
		fmt.Printf("o número %d é par, e também é menor que 10 \n", a)

	case parMaior10(a):
		fmt.Printf("o número %d é par, e também é maior que 10 \n", a)

	default:
		fmt.Printf("o número %d é ímpar \n", a)
	}

}

// ==============================================
// case chamando uma função
func parMenor10(v int) bool {
	if v%2 == 0 {
		return v < 10
	}

	return false

}

func parMaior10(v int) bool {
	if v%2 == 0 {
		return v > 10
	}

	return false
}
