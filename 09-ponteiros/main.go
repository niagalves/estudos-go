package main

import "fmt"

func main() {

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var2)
	var1++

	// ela fica sem ligação uma com a outra, ou seja
	// primeiro print é 10 e o segundo é 11
	// pois tu apenas passou uma copia
	fmt.Println(var1)

	// ponteiro é uma referencia de memoria, um endereço
	// de memoria onde está salvo esse valor
	var var3 int
	var ponteiro *int

	var3 = 100
	ponteiro = &var3

	// vai retornar algo do tipo 0xc000112018
	// que é o endereço da memoria
	fmt.Println(ponteiro)

	// para ver o valor correto, vc faz a
	// desreferenciação
	fmt.Println(*ponteiro, var3)
}
