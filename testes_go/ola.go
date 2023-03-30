package main

import "fmt"

func Hello(resultName string) string {
	return "Olá, " + resultName
}

func WaitHello(waitName string) string {
	return "Olá, " + waitName
}

func main() {
	fmt.Println(Hello("qualquer coisa"))
	fmt.Println(WaitHello("hello"))
}
