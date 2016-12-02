package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123
	fmt.Println(int(y) + x) // converts y from float64 to int; this is a narrowing conversion since
	// you lose all of the values in the decimal places
}
