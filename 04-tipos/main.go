package main

import (
	"errors"
	"fmt"
)

func main() {
	// Existem 4 tipos de numeros inteiros no go
	// int8, int16, int32, int64
	// a diferença de int é a quantidade de bits que cada um suporta

	var numero int16 = 100
	var numero2 int8 = 10

	// quando voce não especifica o int, ele vai usar de acordo
	// com sua arquitetura do computador
	var numero3 int = 8

	fmt.Println(numero, numero2, numero3)

	// uint -> unsygned int
	// var numero4 uint32 = -1000
	var numero4 uint32 = 1000

	fmt.Println(numero4)

	// alias
	// int32 = rune
	var numero5 rune = 99
	fmt.Println(numero5)

	//  byte = uint8
	var numero6 byte = 33
	fmt.Println(numero6)

	// float existe apenas o float32 e o float64
	var numReal float32 = 123.33
	var numReal2 float64 = 21212121221.21

	fmt.Println(numReal, numReal2)

	// Strings
	var str string = "qualquer texto"
	fmt.Println(str)

	// char
	// é o numero da posição que ele tá na tabela
	char := 'B'
	fmt.Println(char)

	// booleano
	var bool1 bool
	fmt.Println(bool1)

	// erro
	// usado em diversas partes das aplicações
	var err error
	fmt.Println(err)

	// existe um pacote erros que é nativo do go
	var err2 error = errors.New("Erro interno")
	fmt.Println(err2)
}
