// Uma interface é um tipo de
// dados abstrato que define um conjunto de métodos.
// diferença struct interface
// struct é usado para representar dados,
// enquanto interface é usada para definir comportamentos.
package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	alture  float64
	largura float64
}

type circulo struct {
	raio float64
}

// interface só tem assinaturas de metodos
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("area da forma é %0.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.alture * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func main() {

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

}
