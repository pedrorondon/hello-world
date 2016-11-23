package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

type Shape interface { // when we declare the Shape interface, it has to have the method on line 17
	// in other words, to define an interface, you just give it a method; anything that has tha method will
	// implement the interface
	area() float64
}

func (s Square) area() float64 { // Square implements the Shape interface; it SATISFIES the Shape interface since it
	// possesses the methods specified by the Shape interface
	return s.side * s.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}
func main() {
	s := Square{10}
	c := Circle{5}
	info(s) // here, the concrete type implementing the Shape interface is a Square
	info(c) // the concrete type implementing the Shape interface is a Circle
}
