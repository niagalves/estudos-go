package main

import "fmt"

func main() {

	numero := 100

	if numero > 15 {
		fmt.Println("maior")
	} else {
		fmt.Println("menor")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("maior que zero")
	}
}
