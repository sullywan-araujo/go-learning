package main

import "fmt"

func Print[T any](v T) {
	fmt.Println(v)
}
