package main

import "fmt"

func main() {
	var x rune = 'a' // rune is an alias for int32
	var y int32 = 'b'
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x)) // converts x from a rune to a string
	fmt.Println(string(y))
}
