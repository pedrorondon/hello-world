package main

import "fmt"

func main() {
	m := make([]string, 1, 25) // slices[] inherently reference a memory address
	fmt.Println(m)             // will print an empty slice[]
	cambio(m)
	fmt.Println(m)
}
func cambio(x []string) { // parameter() receive the slice and assign it to variable, x
	x[0] = "Dro" // code says take the item at position 0 and make it "Dro"
	fmt.Println(x)
}

// when passing a slice[], the address is automatically being passed
// a slice[] is known as a reference type
