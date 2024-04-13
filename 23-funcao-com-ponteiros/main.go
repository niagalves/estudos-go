package main

import "fmt"

func inverteSinal(n int) int {
	return n * -1
}

func inverteSinalComPonteiro(n *int) {
	*n = *n * -1
}

func main() {
	num := 20
	numInvertido := inverteSinal(num)

	fmt.Println(numInvertido)
	fmt.Println(num)

	novoNum := 40
	inverteSinalComPonteiro(&novoNum)
	fmt.Println(novoNum)

}
