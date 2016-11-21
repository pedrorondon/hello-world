package main

import "fmt"

func main() {
	x := 0 // bring x := 0 here narrows the scope; it is no longer at package level like
	// previous example
	increment := func() int { // func expression with an anonymous function of type func() int
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Printf("%T \n", increment)
}

// func expression allows you to have a function inside of a function
