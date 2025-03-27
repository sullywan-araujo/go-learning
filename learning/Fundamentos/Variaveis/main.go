package main

import "fmt"

func main() {

	// Declaração variáveis
	var test string = "testando"
	fmt.Printf("%T", test)

	test1 := map[string]int64{
		"test":  1,
		"seila": 50,
	}
	fmt.Printf("%T", test1)

	var test2 bool
	fmt.Printf("%T", test2)

	// ==================================================
	// INTEGER

	var numberInt8 int8 // suporta entre -128 ~ 127
	numberInt8 = 127
	fmt.Println(numberInt8)

	var numberInt16 int16 // suporta entre -32.768 ~ 32.767
	numberInt16 = -32768
	fmt.Println(numberInt16)

	var numberInt32 int32 // suporta entre -2.147.483.648 ~ 2.147.483.647
	numberInt32 = 2147483647
	fmt.Println(numberInt32)

	var numberInt64 int64 // suporta entre -9.233.372.036.854.775.808 ~ 9.233.372.036.854.775.807
	numberInt64 = -9223372036854775808
	fmt.Println(numberInt64)

	// ==================================================

	// UINT: não aceita números negativos

	var numberUInt8 uint8 // suporta entre 0 ~ 255
	numberUInt8 = 255
	fmt.Println(numberUInt8)

	var numberUInt16 uint16 // suporta entre 0 ~ 65.535
	numberUInt16 = 65535
	fmt.Println(numberUInt16)

	var numberUInt32 uint32 // suporta entre 0 ~ 4.294.967.295
	numberUInt32 = 4294967295
	fmt.Println(numberUInt32)

	var numberUInt64 uint64 // suporta entre 0 ~ 18.446.744.073.709.551.615
	numberUInt64 = 18446744073709551615
	fmt.Println(numberUInt64)

	// ==================================================

	// FLOAT
	// quando iniciamos uma variável com um npumero decimal e não definimos o tipo, por padrão ele será float32;

	// ==================================================
	//ANY
	test3 := map[string]interface{}{
		"test123": "test456",
		"test4":   false,
		"test5":   []int64{1, 2, 3, 4, 5},
	}

	fmt.Println(test3)
	fmt.Printf("%T", test3)

	// ==================================================
	//STRUCT

	var user User = User{
		Name:     "Sullywan",
		LastName: "Araujo",
		Age:      40,
	}

	fmt.Println(user)

}

// Variáveis/Funções públicas e privadas
// o que difere uma variável/função pública de uma variável/função privada, é a letra inicial da variável;

func FuncPub() {

}

func funcPriv() {

}

// ==================================================
//STRUCT

type User struct {
	Name     string
	LastName string
	Age      int
}
