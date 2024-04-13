// panic mata a execução do programa
// recover recupera a aplicação
// se sua aplicação não exitver em panic, nào irá acontecer
// nada e o recover irá ignorar
// panic não é a melhor alternativa para capturar erro
// ele é diferente do erro, pois tem outros metodos
package main

import "fmt"

func recuperarExec() {
	fmt.Println("tentativa de recuperar")

	if r := recover(); r != nil {
		fmt.Println("recuperado com sucesso")
	}
}

func calcula(n1 float64, n2 float64) bool {
	defer recuperarExec()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é 6")
}

func main() {
	fmt.Println(calcula(6, 6))
	fmt.Println("pós execução")
}
