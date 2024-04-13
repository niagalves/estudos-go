package auxiliar

import "fmt"

// Letra maiuscula permite que o arquivo seja importado
// Go não é orientado a objetos então ão existe private, public, etc
// Se for maiusculo ela permite ser importada para outros pacotes
// caso seja minuscula ela é privada apenas nele.
// Quando o arquivo for importado, é bom ter um comentário tipo
// esse para explicar o que a função faz
func Escrever() {
	fmt.Println("Escrevendo pacote auxiliar")

	// posso usar ela dentro do próprio pacote
	escrever2()
	verificaEmail()
}
