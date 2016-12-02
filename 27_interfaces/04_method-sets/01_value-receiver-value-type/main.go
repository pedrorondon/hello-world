package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c circle) area() float64 { // c cirle is a value receiver which can take either a value
	// or a pointer
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println("Area is", z.area())
}

func main() {
	x := circle{5}
	info(x) // value passed to the value receiver c circle
}
