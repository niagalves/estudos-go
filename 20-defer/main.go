// ela adia a execuÇào de um determinado bloco de código
package main

import "fmt"

func func1() {
	fmt.Println("exec 1")
}

func func2() {
	fmt.Println("exec 2")

}

func main() {

	defer func1()
	func2()
}
