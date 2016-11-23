package main

import "fmt"

type Square struct { // create a user defined type and call it Square; the underlying type is a struct
	side float64
}

func (s Square) area() float64 { // attach a method to the type we created, Square
	return s.side * s.side
}

type Shape interface { // if you want to implement the Shape interface, you need to have a method with
	area() float64 // <------ this signature; as seen in line 9, type Square has a method that has that
	// exact signature; therefore, Square implements the Shape interface; now, for anything
	// in this example, the method needs to be called area, have no parameters, and returns a float64

	// Bill Kennedy: Interfaces are types that just declare behavior.  This behavior is never implemented
	// by the interface directly, but instead by user-defined types via methods.
}

func info(z Shape) { // create a new function that can take a type as an argument; here it is taking type Shape;
	// you can pass anything that implements type Shape into this function; you can also pass in Square for anything
	// tha asks for Shape because Square implements the Shape interface
	fmt.Println(z) // prints the value of type Square
	fmt.Printf("%T \n", z)
	fmt.Println(z.area())
	fmt.Printf("%T \n", z.area())
}
func main() {
	s := Square{10}
	info(s) // here you are calling function info and passing in argument s; s is a variable of type Square;
	// we can pass s into function info which is expecting Shape but will accept Square since Square is a Shape
	// because it implements the Shape interface by having a matching method
}
