package main

import "fmt"

func main() {
	remind := 7.24
	fmt.Printf("%T\n", remind)
	fmt.Printf("%T\n", int(remind))
}

// assertion with variable.(type) on the right side; the variable will be asserted to the type
// named in parens to the right of the dot.
// casting with type(variable) on the left side; the variable in parens() will be cast
// to the type named on the left
