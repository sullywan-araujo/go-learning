package funcaoretornonomeado

func Transferir(origem, destino string, saldoOrigem, valor float64) (novoSaldoOrigem float64, sucesso bool) {
	if saldoOrigem > valor {
		novoSaldoOrigem = saldoOrigem - valor
		sucesso = true
	} else {
		novoSaldoOrigem = saldoOrigem
		sucesso = false
	}

	return
}
