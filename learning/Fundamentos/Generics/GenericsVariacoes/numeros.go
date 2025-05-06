package main

type Numerico interface {
	~int | ~float64
}

func Soma[T Numerico](a, b T) T {
	return a + b
}
