// Em Go, sync.WaitGroup é uma estrutura útil
// para controlar a concorrência em programas que
// utilizam goroutines. Goroutines são como threads leves,
// permitindo que você execute várias tarefas de forma concorrente.

// O sync.WaitGroup é usado para esperar até que todas as
// goroutines que estão sendo executadas terminem antes de
// continuar a execução do programa principal.
// Isso é especialmente útil quando você precisa
// esperar que várias goroutines completem uma
// tarefa antes de prosseguir para a próxima etapa do seu programa.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em go!")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
