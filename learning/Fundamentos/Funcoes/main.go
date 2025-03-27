package main

import (
	"errors"
	"fmt"

	fr "github.com/learning/Fundamentos/Funcoes/FuncaoComRetorno"
	fmr "github.com/learning/Fundamentos/Funcoes/FuncaoMultiRetorno"
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

	//=============================================
	novoSaldo, sucesso := fmr.Sacar(saldo, 1950.00)

	if sucesso {
		fmt.Println("Saque realizado com sucesso!")
	} else {
		fmt.Printf("Falha ao sacar, saldo insulficiente: %2.f\n", saldo)
	}

	fmt.Printf("Novo saldo atualizado: %2.f\n", novoSaldo)

}

func test() (string, error) {
	return "", errors.New("test")
}
