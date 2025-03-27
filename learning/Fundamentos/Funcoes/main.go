package main

import (
	"errors"
	"fmt"

	fr "github.com/learning/Fundamentos/Funcoes/FuncaoComRetorno"
	fs "github.com/learning/Fundamentos/Funcoes/FuncaoSemRetorno"
)

func main() {

	variavel, err := test()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(variavel)

	//=============================================
	nomeCliente := "Joao"
	saldoAtual := 1550.00

	fs.ExibirSaldo(nomeCliente, saldoAtual)

	//=============================================
	saldo := fr.CalcularSaldo(saldoAtual, 500.00)

	fmt.Println("Saldo atualizado:", saldo)

}

func test() (string, error) {
	return "", errors.New("test")
}
