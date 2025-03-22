package main

import (
	"fmt"
	"sort"
)

func main() {
	/*nums := []int{10, 20, 30}

	for i := 0; i < len(nums); i++ {
		nums = append(nums, nums[i]+10) // a soma está sendo feita sobre o valor que está na posição do índice.
		fmt.Println("Valor de i: ", i)
		fmt.Println("Valor de nums", nums)
	}*/
	// a expressão que determina o fim do looping é analizado a cada iteração e como está sendo adicionado novos itens ao slice o looping acaba sendo infinito.

	//=====================================================
	// for com range
	// Nesse caso não entra em um looping infinito, pois a função 'range' cria uma cópia do slice para fazer a condição,
	// ou seja, no caso abaixo temos 3 elementos e mesmo que seja acrescentado elementos ao slice, a clausula 'range' não considera os elementos
	// na verificação da condição da cláusula;

	nums := []int{1, 2, 3}

	for i, v := range nums {
		fmt.Println("Índice:", i, "Valor:", v)
		nums = append(nums, v+10) // Adiciona novos elementos
	}

	fmt.Println("Final:", nums)

	//=====================================================
	// iteração ordenada crescente

	// map de pagamentos (ID da transação / valor)
	pagamentos := map[int]float64{
		1001: 234.00,
		234:  343.45,
		102:  550.00,
		457:  490.00,
	}

	// criando um slice para armazenar os ID's das transações;
	ids := make([]int, 0, len(pagamentos))

	// preenchendo o slice 'ids' com as chaves do map
	for id := range pagamentos {
		ids = append(ids, id)
	}

	// ordenando os IDs das transações
	sort.Ints(ids)

	// iterando sobre os IDs ordenados e acessando os valores do map
	for _, id := range ids {
		fmt.Printf("ID: %d | valor: R$ %.2f\n", id, pagamentos[id])
	}

	//=====================================================
	// iteração ordenada decrescente
	pag := map[int]float64{
		10: 304.50,
		5:  39.00,
		45: 493.00,
	}

	trs := make([]int, 0, len(pag))

	for tr := range pag {
		trs = append(trs, tr)
	}

	sort.Slice(trs, func(i, j int) bool {
		return pag[trs[i]] < pag[trs[j]]
	})

	for _, tr := range trs {
		fmt.Printf("ID da transação: %d | valor: R$ %.2f\n", tr, pag[tr])
	}

}
