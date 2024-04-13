package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// funções podem ter mais que um retorno
func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	soma := somar(10, 21)

	fmt.Println(soma)

	var f = func(txt string) string {
		return txt
	}

	resultado := f("função f")
	fmt.Println(resultado)

	resultadoSoma, resultadoSub := calculos(10, 4)
	fmt.Println(resultadoSoma, resultadoSub)

	// Não quero um dos resultados basta utilizar _
	_, res := calculos(10, 4)
	fmt.Println(res)
}
