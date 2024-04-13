package main

import "fmt"

func main() {

	// maps é uma outro estrutura
	// para guardar dados
	// uma das caracteristicas do map
	// é manter os tipos coerentes
	usuario := map[string]string{
		"nome":      "Niag",
		"sobrenome": "Alves",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joao",
			"ultimo":   "niag",
		},
		"faculdade": {
			"curso":     "ADS",
			"matricula": "fac21111",
		},
	}

	fmt.Println(usuario2)

	// deletar um map
	delete(usuario2, "faculdade")
	fmt.Println(usuario2)

	// adicionar um novo
	usuario2["signo"] = map[string]string{
		"nome": "Touro",
	}

	fmt.Println(usuario2)
}
