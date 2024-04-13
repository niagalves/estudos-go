// ele é uma coleção de campos de dados
// cada campo tem um nome e um tipo
// são agrupados em uma unica unidade

package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	// existem três maneiras de criar uma struct
	var u usuario
	u.nome = "Niag"
	u.idade = 31
	fmt.Println(u)

	end := endereco{"Rua da praia", 10}

	u2 := usuario{"Dav", 21, end}
	fmt.Println(u2)

	// quando vc não sabe os dados completos ainda,
	// pode utilizar desse modo
	u3 := usuario{nome: "Jo"}
	fmt.Println(u3)
}
