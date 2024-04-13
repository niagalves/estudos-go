// são parecidos com funções
package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvar %s \n", u.nome)
}

func main() {
	usuario := usuario{"Niag", 20}
	fmt.Println(usuario)

	usuario.salvar()

}
