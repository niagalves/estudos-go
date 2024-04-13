package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "Dom"
	case 2:
		return "Seg"
	case 3:
		return "Ter"
	case 4:
		return "Qua"
	case 5:
		return "Qui"
	case 6:
		return "Sex"
	case 7:
		return "Sab"
	default:
		return "inválido"
	}

}

func main() {

	result := diaSemana(10)
	fmt.Println(result)

}
