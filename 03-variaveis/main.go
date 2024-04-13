// go é uma linguagem fortemente tipada
// então é necessário saber os seus tipos, etc
package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	// desse modo ele ja identifica o tipo dela
	// tecnica chamada de inferencia de tipo
	variavel2 := "variavel 2"

	// Criar mais de uma variavel ao mesmo tempo
	var (
		variavel3 string
		variavel4 string
	)

	variavel5, variavel6 := "Variavel 3", "Variavel4"

	fmt.Println(variavel1, variavel2, variavel3, variavel4, variavel5, variavel6)

	// invertendo o valor de uma variavel
	variavel5, variavel6 = variavel6, variavel5
}
