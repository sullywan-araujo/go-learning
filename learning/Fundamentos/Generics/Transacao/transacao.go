package transacao

type Transacao[T float64 | int] struct {
	ID    string
	Valor T
}

func TransacaoMaiorValor[T float64 | int](a, b Transacao[T]) Transacao[T] {
	if a.Valor > b.Valor {
		return a
	}

	return b
}
