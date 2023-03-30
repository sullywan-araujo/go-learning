package main

import "fmt"

const prefixOlaPt = "Olá, "

func Hello(resultName string) string {
	if resultName == "" {
		resultName = "Mundo"
	}
	return prefixOlaPt + resultName
}

func WaitHello(waitName string) string {
	return prefixOlaPt + waitName
}

func main() {
	fmt.Println(Hello("qualquer coisa"))
	fmt.Println(WaitHello("hello"))
}
