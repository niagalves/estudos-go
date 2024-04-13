package main

import (
	"fmt"
	"reflect"
)

func main() {

	// arrays e slices são duas estruturas parecidas
	// mas temos que saber a hora certa de utilizar elas

	var arr1 [5]int
	fmt.Println(arr1)

	// acessar posicao
	fmt.Println(arr1[1])

	arr2 := [5]string{"pos 1", "pos 2", "pos 3", "pos 4", "pos 5"}
	fmt.Println(arr2)

	// assim ele fixa com a quantidade de itens
	// de acordo com o que tu passar
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	// para criar um slice é bem parecido com o array
	// esse cara é muito usado em go
	// ele não limita o tamanho das coisas
	// slice não é array, é apenas parecido
	// é uma fatia de um array
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	// a diferença está no valor setado
	// aqui retorna [] para slice e [5] para array
	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(arr3))

	// temos uma função no slice chamada append
	// para adicionar um novo item no slice
	slice = append(slice, 100)
	fmt.Println(slice)

	// pegar de um indice até o outro
	slice2 := arr2[1:3]
	fmt.Println(slice2)
}
