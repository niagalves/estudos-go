package main

import "fmt"

func calculos(n1 int, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	fmt.Println(calculos(20, 15))

}
