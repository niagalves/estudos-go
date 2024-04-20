// ponto forte do go é a concorrencia
// por ser muito fácil de utilizar
// existe uma diferença entre
// concorrencia e paralelismo

// paralelismo é quando duas ou mais tarefas
// estao atuando ao mesmo tempo

package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo!")
	escrever("Programando em go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
