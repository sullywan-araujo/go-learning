package main

func Contem[T comparable](lista []T, item T) bool {
	for _, v := range lista {
		if v == item {
			return true
		}
	}

	return false
}
