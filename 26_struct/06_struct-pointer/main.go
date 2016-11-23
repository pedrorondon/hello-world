package main

import "fmt"

type person struct { // user defined type
	nombre string
	edad   int
}

func main() {
	p1 := &person{"Wolverine", 199}
	fmt.Println(p1)
	fmt.Printf("%T \n", p1) // of type pointer to a function
	fmt.Println(p1.nombre)  // you can pull the values from a pointer
	fmt.Println(p1.edad)
}
