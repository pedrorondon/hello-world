package main

import "fmt"

func main() {
	lett := rune("A"[0]) // string to rune conversion; string is wrapped in double quotes ""
	// this line of code says take the first position, position 0, of the string, "A", and turn it
	// into a rune
	fmt.Println(lett)
}

// string is a collection of runes and runes are bytes; therefore, a string is a slice of bytes
// rune is an alias for int32
