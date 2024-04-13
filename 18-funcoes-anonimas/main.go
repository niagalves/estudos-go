package main

import "fmt"

// são funcões que não possuem nomes e são
// executadas usando () ao final
// ela executa após rodar de imediato
func main() {

	func() {
		fmt.Println("Olá")
	}()

	func(txt string) {
		fmt.Println(txt)
	}("Olá parametro")
}
