package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		time.Sleep(time.Second)
		fmt.Println("incrementando", i)
		i++
	}

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println("incrementando", j)

	}

	nomes := [3]string{"João", "Davi", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
}
