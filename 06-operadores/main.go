package main

import "fmt"

func main() {
	// aritméticos
	// + - / * %

	soma := 1 + 2
	sub := 1 - 2
	div := 10 / 3
	mult := 3 * 3
	resto := 10 % 2

	fmt.Println(soma, sub, div, resto, mult)

	// Não é possivel fazer operações aritméticas
	// com tipos diferentes

	var num1 int16 = 10
	var num2 int16 = 25

	soma2 := num1 + num2
	fmt.Println(soma2)

	// Atribuição
	var variavel string = "string"
	variavel2 := "string2"
	fmt.Println(variavel, variavel2)

	// operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)

	// operadores logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro || falso)

	// unários
	numero := 10
	numero++
	numero += 100
	numero--
	numero *= 1
	numero /= 2
	numero %= 2
	fmt.Println(numero)
}
