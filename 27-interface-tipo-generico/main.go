package main

import "fmt"

// desse modo pode ser tudo
// parece muito com o any do TS
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(1)
	generica(true)
}
