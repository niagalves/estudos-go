package main

import "fmt"

func main() {
	// Se utiliza a funcão make
	slice := make([]float32, 10, 15)
	fmt.Println(slice)
	fmt.Println(len(slice)) // tamanho
	fmt.Println(cap(slice)) // capacidade -> opcional

	// o go quando vê que seu array irá estourar
	// ele pega e dobra de tamanho
	// 10 + 10 = 20
	// é assim que funciona o array interno
}
