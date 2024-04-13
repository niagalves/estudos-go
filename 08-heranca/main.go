// é o que o go chega mais proximo do POO
package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

// apenas adiciona o type sem tipar,
// assim ele consegue acessar os dados do struct
type estudante struct {
	pessoa
	matricula uint64
}

func main() {
	p := pessoa{"Niag", 31}
	e := estudante{p, 888}

	fmt.Println(p, e)

}
