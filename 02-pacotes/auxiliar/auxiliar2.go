package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func escrever2() {
	fmt.Println("Escrevendo 2")
}

// Para instalar pacotes externos nós executamos o comando
// go get url do pacote exemplo:
// go get go get github.com/badoux/checkmail
func verificaEmail() {
	err := checkmail.ValidateFormat("dev@gmail.com")

	fmt.Println(err)
}

// Para remover os pacotes não utilizados do go mod
// execute o comando go mod tidy
