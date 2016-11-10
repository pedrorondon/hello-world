package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a, "a")
	fmt.Println(&a, "b")
	var b *int = &a
	fmt.Println(b, "c")
}

// the code above makes b a pointer to the memory address where an integer is a stored
// b is of type "int pointer"
// *int -- the * is part of the type -- b is of type *int
