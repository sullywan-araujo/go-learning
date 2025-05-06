package main

type Caixa[T any] struct {
	Valor T
}

func (c Caixa[T]) Get() T {
	return c.Valor
}
