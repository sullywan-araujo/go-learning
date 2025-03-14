package main

import "fmt"

func main() {
	//IF
	if 2 > 1 {
		fmt.Println("true")
	}

	// IF / ELSE
	if 1 > 1 {
		fmt.Println("Equals")
	} else {
		fmt.Println("Different")
	}

	// IF / ELSE IF
	if 2 < 1 {
		fmt.Println("falso")
	} else if 5 < 1 {
		fmt.Println("falso2")
	} else {
		fmt.Println("verdadeiro")
	}
}
